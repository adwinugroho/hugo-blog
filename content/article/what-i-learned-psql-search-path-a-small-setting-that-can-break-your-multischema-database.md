---
title: "What I Learned Psql Search Path a Small Setting That Can Break Your Multischema Database"
date: 2026-10-03T17:51:09+08:00
draft: true
---
### Introduction
As a Backend Engineer we used to think simple and easy way to manage our database. We often use ORM/Library 3rd party to interact with our database, that's fine that's totally fine. But it will be a nightmare when we really don't understand or know what happen behind the scene and suddenly something goes wrong. In this article, I want to share my experience about multischema database in PostgreSQL and "search_path".

1. What is multischema database
Having a huge system like chat application on omnichannel platform need a robust database and great scalable architecture. We have 2 choices, Database per tenant or single database with multiple schema. We go with the second option which is single database with multiple schema.

Example scenario, we have 3 tenants : A, B, and C. If we use database per tenant, we need to create 3 database. If we use single database with multiple schema, we need to create 3 schema.
 
2. What is "search_path"`
### Architecture
### The Incident
1. Why did this happen?
2. The debugging process
3. What we changed?
### Conclusion and Lessons Learned
1. Why it's crucial to understand the tools you're using
2. 
