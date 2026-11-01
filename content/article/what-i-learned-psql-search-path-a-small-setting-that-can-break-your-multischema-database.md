---
title: "What I Learned: How a Small `search_path` Setting Can Break Your Multi-Schema Database"
date: 2026-10-03T17:51:09+08:00
draft: false

categories: ["TIL", "WIL", "Incident", "Research"]
tags: ["architecture", "system design", "sql", "postgreSQL", "golang"]
toc: false
author: "adwin"
---

### Introduction
As Backend Engineers, we often look for the simplest and easiest way to manage our databases. We frequently use ORMs or third-party libraries to interact with our database, and that's fine. That's totally fine. But it becomes a nightmare when we don't truly understand what happens behind the scenes, and suddenly, something goes horribly wrong. 

In this article, I want to share my real-world experience dealing with a multi-schema database in PostgreSQL and the infamous `search_path`.

### 1. What is a Multi-Schema Database?
Running a massive system like an omnichannel chat application requires a robust database and a highly scalable architecture. We basically had two choices: Database-per-tenant, or a single database with multiple schemas. We went with the second option: a single database with multiple schemas.

For example, if we have 3 tenants (A, B, and C), a database-per-tenant approach means creating 3 separate databases. A multi-schema approach means creating 3 schemas within one database. Another option is table partitioning, where a table like `orders` is automatically broken down by PostgreSQL into child tables like `order_2026`, `order_2025`, `order_2024`, etc. Although table partitioning is more efficient for executing huge queries, we chose the multi-schema architecture for this specific use case.

![Example of omnichannel chat architecture](/image/content/article/what-i-learned-psql-search-path-a-small-setting-that-can-break-your-multischema-database/image-1.png)

The main reason we picked multi-schema was **isolation**. We believed it was the best solution for this case. We also managed our database migrations using `go-migrate`. If we have a thousand or more tenants, it would be difficult to migrate them otherwise. Plus, with this architecture, we can simply select tables for each store without worrying about data leaking to another store.

### 2. What is `search_path`?
Since we chose multi-schema, we knew PostgreSQL has a built-in feature called `search_path`. What is it? It’s a mechanism that allows PostgreSQL to resolve unqualified SQL object names, selecting the object referenced without explicitly specifying its schema [^1]. 

For example, if we have a tenant named "txdawinganteng" and we want to query the `buyer` table. Instead of doing this:

```sql
SELECT id, name FROM "txdawinganteng".buyer WHERE channel_id = 4;
```

(Assuming search_path is already set to txdawinganteng)

### The Incident
### 1. Why did this happen?
Once we understood how `search_path` worked, we planned to implement it right in our middleware. At the first layer, we thought we knew what we were doing, so we wrote this code:

```go
func DBMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tenantID := getTenantFromContext(r.Context())
        schemaName := fmt.Sprintf("tenant_%s", tenantID)
        
        _, err := db.Exec(r.Context(), "SET search_path TO "+schemaName)
        if err != nil {
            http.Error(w, "DB Error", http.StatusInternalServerError)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

We deployed it to staging. Nothing happened. It just worked as planned. Everything was fine. Our QA team even confirmed that this implementation didn't introduce any defects or bugs. Of course, after all tests passed, we planned a production deployment at midnight. We deployed this implementation with a few minor fixes. Regularly, after deployment, we tested as usual. The developers, including me and all my teammates, ran manual tests. We sent chats to all the stores we had, and everything seemed fine.

The next morning, I went to the office and was absolutely shocked. There were many reports and tickets from the Customer Service team. They reported a **critical bug: chats from Tenant A were going to the store of Tenant B. Hahaha.
Not just that. Yes, the app was loading slowly, and many timeouts were happening.** Hmm, I was so confused. Why is this happening? What went wrong? Did I make a mistake? We immediately had to call the DevOps team to revert the image (pod).

### 2. The Debugging Process
I had to dive into the debugging process to tackle this critical bug, reverting the code I deployed the night before one by one. However, I didn't see the problem in the minor fixes because we hadn't changed how we connected to the database.
I eventually realized that `search_path` was the reason why many chats were missing or going to the wrong tenant. But the real question is: *why?* On top of that, why was my app suddenly so slow? Hmm, interesting. I asked our QA tester to reproduce this error, but somehow, on staging, the error didn't happen at all.

### 3. What We Changed (The Root Cause)
After an hour of debugging, I realized the reproduction environment was not the same. The data on staging was much smaller than in production. In staging, we only had 3-5 tenants with 2-5 stores each. Whereas in production, we have massive data; just one tenant can have 25 stores or even more, hahaha.

I decided to test with `k6` using dummy tenants and stores. The marketplaces like Shopee and TikTok already have proper webhooks, and yes, the error appeared. It was exactly the same as in production.

As far as I know, the main problem was initiating `SET search_path TO` after we got a connection from the pool. Even worse, we put it in the middleware. This causes a race condition and massive overhead on DB connections. The `pgx` connection pooling is simply not compatible with dynamic search_path manipulation per request, leading to a classic connection pooling trap.

We all know connection pooling is a strategy to manage database connections for multiple requests. By changing the session state mid-pool, we were essentially poisoning the connections for the next user.

### Conclusion and Lessons Learned
### 1. Why it's crucial to understand the tools you're using
We all know we can't survive without using a tech stack, whether it's open source, freemium, premium, or whatever you name it: Redis, PostgreSQL, jwt.io, Supabase, Kafka, MongoDB, etc. Maybe big companies like Google, Meta, or Apple use built-in libraries, or if they use open source, they customize it heavily.

We can't fully control what we use. Yeah, you're right, we can make an issue or submit PRs, but how much can we really do to fix that tech ourselves if we must customize it? All that matters is doing what we can: learning and deep-diving into the tech stack we use.

It cannot just be: "Oh yeah, GORM is easy to use, we're gonna use it anyway!" or "Oh, I love Elastic APM, I just call a function and it's ready to serve!" No, it is not like that! So, what is the actual correct implementation?
In my case, I completely removed the `SET search_path logic` from the middleware. Instead, I made the schema explicit by passing the **schemaName** as a parameter to every repository function. This guarantees that the query is stateless and immune to connection pool pollution.

```go
func (r *MessageRepo) GetMessages(ctx context.Context, schemaName string, tenantID string) ([]Message, error) {
    // IMPORTANT: Always sanitize schemaName to prevent SQL Injection!
    safeSchema := sanitizeSchemaName(schemaName) 
    
    query := fmt.Sprintf("SELECT id, content FROM %s.messages WHERE tenant_id = $1", safeSchema)
    
    rows, err := r.db.Query(ctx, query, tenantID)
}
```

You might be wondering, "What about PGX's AfterConnect hook?" Yes, you can configure the connection pool to automatically run `SET search_path` whenever a new physical connection is established. However, this is only safe if that specific pool is strictly dedicated to a single tenant. If you share that pool across multiple tenants dynamically, you are back to square one with race conditions. For a system with thousands of dynamic tenants, managing thousands of separate connection pools is a resource nightmare.

When should we actually use search_path? Honestly, its best use cases are highly restricted:
- Local Development: When you are debugging on your localhost and don't want to type tenant_123.table_name a hundred times.
- Isolated Cron Jobs or Maintenance Scripts: For background jobs that run weekly or monthly (e.g., generating a specific tenant's monthly report), where the script explicitly sets the search_path for its own dedicated, short-lived connection, completely isolated from the main web application's connection pool.

### 2. Lesson Learned
In the age of AI, the line between a "code monkey" and a real Software Engineer is being tested more than ever. Maybe you think LeetCode, HackerRank, or competitive programming isn't reliable. Maybe you are right, maybe you are wrong.
But we should learn how to think and how to implement, not just chase results that look extraordinary, perform well, or seem unbelievable on the surface. Because if you work in a big company and build vital services like fintech, payment gateways, banking apps, or just an app with millions of users, milliseconds matter. Every hour, every minute, even every second is important. Users' money must be safe, and the users are the core value of the business.

Well, what if I am just an "Indihacker"? We should ship more products, right? All that matters is the client, and the client only knows if the program works well or not. Yeah, you're right. Until you ship a premium product with payments, and you suddenly realize the payment gateway doesn't work well. Hopefully, we are not the ones causing that. Whatever you are—an Indihacker, a Vibecoder, a Software Engineer, or a a Fullstack Developer—please learn what you use, and stay curious!

---
References:
[^1]: PostgreSQL: an introduction to search_path
