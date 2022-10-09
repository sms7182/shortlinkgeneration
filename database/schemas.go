package database

const createSchema = `
CREATE TABLE IF NOT EXISTS shorturllinks
(
	id SERIAL PRIMARY KEY,
	url TEXT,
	shortlink TEXT,
	count INTEGER
)
`

var insertShortUrlLink = `
INSERT INTO shorturllinks(url, shortlink, count) VALUES($1,$2,$3) RETURNING id
`
var getLastShortLink = `
select max(sul.count) from public.shorturllinks sul
`

var getActualUrl = `
select sul.url  from public.shorturllinks sul where sul."shortlink" = '%v'
`
