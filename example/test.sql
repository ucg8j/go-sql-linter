WITH bad_cte AS (
  SeLeCT 
    a,   
    b,  


  from `project.dataset.table` a
),



good_cte AS (
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
d,
FROM bad_cte
LEFT JOIN good_cte USING (a)
