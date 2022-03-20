Q:  How to choose DB ?
A:  It is important to choose speacific DB because it impacts your non fucntional reuqiremtns on how your design scale.
    It doesn't impacts your fucntional reuqiremtns but impacts lot your non fucntional reuqiremtns.
    1. Structure of the data we have in question. Does we have Structured data or completely non structured data
    2. Query pattern.
    3. Amout of scale you want in question.

    Blob Storage: For media and files. ex: amazon s3
    CDN: Geographically districute data 
    Text searching : Users to search on text. Elastic Search/open Search and Solr [Lucene] + Support fuzzy search
    
    Metrics data : Aps pushing metrics to DB. Time series DB. exmaple : InfluxDB, OpenTSDB
    We enerally dont do random read or update on this kind of data.
    Queried mostly for bulk, like last 24Hrs etc.

    Analytics on all Txns :- Datawarehosue, hadoop etc used for offlice reporting.
    

____________________________________________________________________________________________________________________________

### Requirements:
    1. HA, Low latency, durability, consistency and reliability.

    2. API : we should inlcude API's before HLD
    3. We may refrain from adding cache in HLD. Its more about scalability and can come later during discussion.
    4. Queues can come in HLD as they are building blocks of HLD.
    