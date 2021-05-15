{{
	config(
		tags=[
			"pii=false"
		],
	)
}}

WITH bad_cte AS (
  -- ❌ A poorly formatted piece of SQL
  SeLeCT 
    a,   
    b,  


  from `project.dataset.table` a
),



good_cte AS (
  -- ✅ Better
  SELECT
  a,
  c,
  d

  FROM `project.dataset.table` b
)



SELECT
a,
b,
c,
d, -- ❌ trailing comma ❌ newline after select block
FROM bad_cte -- ❌ newline after from block
LEFT JOIN good_cte USING (a)
