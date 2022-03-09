## Crawler

The function starts crawling from a random wiki page and stores it in a leveldb database (which is a key-value store same as rocksdb, rocksdb does not have support for golang as of now)

## Search

In the main func the second call is to search a keyword across the database. It returns the url's containing those links.

`Note: There is a small problem -- Wikipedia blocks the IP address after crawling a few links. This can be rectified by using vpn's on cloud or using something along the lines of DBpedia which is more structured and easy to query. (However it has it's own querying issues that I faced in another project)`

If The IP is blocked to test the search we can comment the `crawler()` function call in main. 