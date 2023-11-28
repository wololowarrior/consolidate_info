package queries

const Insert = `
INSERT INTO Contact(phonenumber, email, linkPrecedence,ipv4)  VALUES($1, $2,$3,$4)
returning id  
-- INSERT INTO contact(phonenumber, email, linkedid, linkprecedence, created_at, updated_at)  VALUES($1, $2, $3, $4, $5, $6)
-- returning id                                                 
`

const GetPrimary = `
SELECT id, email, phonenumber, linkedID, linkPrecedence
FROM Contact
WHERE ipv4=$1 and linkPrecedence='primary'
`

const UpdateLinkedID = `
UPDATE Contact
	set linkedID = $1
	WHERE id=$2
`

const Get = `
SELECT id, email, phonenumber, linkedID, ipv4, linkPrecedence
FROM Contact
WHERE id=$1
`
