---
title: "What I Learned - Psql Search Path a Small Setting That Can Break Your Multischema Database"
date: 2026-10-03T17:51:09+08:00
draft: true

categories: ["TIL", "WIL", "Incident", "Research"]
tags: ["architecture",  "system design", "sql", "postgreSQL"]
toc: false
author: "adwin"
---
### Introduction
As a Backend Engineer we used to think simple and easy way to manage our database. We often use ORM/Library 3rd party to interact with our database, that's fine that's totally fine. But it will be a nightmare when we really don't understand or know what happen behind the scene and suddenly something goes wrong. In this article, I want to share my experience about multischema database in PostgreSQL and "search_path".

1. What is multischema database
Having a huge system like chat application on omnichannel platform need a robust database and great scalable architecture. We have 2 choices, Database per tenant or single database with multiple schema. We go with the second option which is single database with multiple schema.

Example scenario, we have 3 tenants : A, B, and C. If we use database per tenant, we need to create 3 database. If we use single database with multiple schema, we need to create 3 schemas. Other option is have table partitioning, which led table to have many child in example we have table *orders*, PSQL automatically breakdown to these table *order_2026*, *order_2025*, *order_2024* etc. although table partitioning more efficient to execute huge query, we have to choose multi-schema with this architecture

![Example of omnichannel chat architecture](/image/content/article/what-i-learned-psql-search-path-a-small-setting-that-can-break-your-multischema-database/image-1.png)

The main reason why we picked multi-schema is isolation, the best solution at this case, we think is multi-schema. We also make database migrations with go-migrate because if we have a thousand or more tenants it will be difficult to migrate and then at the certain things we can just can select table on each store without worry about data leak to another store.
 
2. What is "search_path"`
We had to choose multi-schema and we know postgreSQL have a built in function called **search_path**. What is **search_path**? This func is a mechanism that allows postgreSQL to resolve unqualified SQL object names, to select object referenced without excplicitly specifying their schema [^1]. In example we have tenant name "txdawinganteng" and querying a select query on table buyer. Despite to do this:
```sql
SELECT id, name FROM "txdawinganteng".buyer WHERE channel_id = 4;
```

We can just do this:
```sql
SELECT id, name FROM buyer WHERE channel_id = 4;
```

### The Incident
1. Why did this happen?
After that, we just know how to **search_path** works and we plan to implement that function at middleware. Eventually we just know at the first layer how to **search_path** works we still execute with this code
```Go
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
Then, we just deployed on staging, nothing happen, it just work as we planned. Everything fine, QA tester at my team too confirmed this implement doesn't defect or have a bug. Ofcourse after all test, we planning to do deployment on production at midnight, we deploy this implement with a few issue and minor fix. Regularly after deploy we test as usuall, include the developer me and all my teammates do the test. We send chat at all store we have and everything fine.

On morning I go to my office I just shocked, many reports and tickets from Customer Service team, they reported the chat have critical bug, chat at tenant A go to store on tenant B hahaha, not just that, yes, the app have a slowly loading many timeout happen. Hmm I just confused why this is happen? What went wrong? Do I make a mistake? Then we have to call devops team to revert the **image** (pod). 

2. The debugging process
So I have to do debugging process to tackle this critical bug, revert one by one code I deployed last night, however I don't see the problem is on the minor fix because we not change how to we connect I just know the **search_path** is the reason of why many chat miss and go on wrong tenant. But the real question is why? On top of that, why my app taking too slow? Hmmm it interesting, I tell to my QA tester to reproduce this error but somehow on staging this error doesn't happen.

3. What we changed?
However after an hour debugging, I think the reproduce is not same, because the data on staging is less more than production, in staging we only have 3-5 tenant with 2-5 store for each tenant whereas in production we have large data, it only 1 tenant can have 25 store even more hahaha. I decide to test with k6 with dummy tenant and store. The marketplace like shopee and tiktok already had a proper webhook and yes the error appear its same as production. As far as I know, the main problem it's on **search_path TO** initiate after we got connection even worst it put on middleware. It causes an race condition and overheard on DB connection because on **pgx pooling** not compitible with the dyanamic **search_path** and led to **connection pooling trap**. We all know the connection pooling is a strategy to manage database connections for multiple requests. 

### Conclusion and Lessons Learned
1. Why it's crucial to understand the tools you're using
2.

References:
[^1]: https://wbarillon.medium.com/search-path-postgresqls-quality-of-life-feature-de3ff5055a17
