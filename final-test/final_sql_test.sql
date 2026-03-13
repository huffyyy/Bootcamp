-- a
select *
from T_POLICY as p
inner join T_CLIENT as c on p.CLIENT_NUMBER = c.CLIENT_NUMBER
where p.POLICY_SUBMIT_DATE > '2018-01-15'
and month(c.BIRTH_DATE) = 9

-- b
select *
from T_POLICY as p
inner join T_AGENT as a on p.AGENT_CODE = a.AGENT_CODE
where p.POLICY_STATUS = 'INFORCE'
and a.AGENT_OFFICE = 'JAKARTA'

-- c
select (p.COMMISSION / p.PREMIUM) * 100 as BASIC_COMMISSION
from T_AGENT as a
inner join T_POLICY as p on a.AGENT_CODE = p.AGENT_CODE

-- d
update T_POLICY
set POLICY_DUE_DATE = last_day(date_add(POLICY_SUBMIT_DATE, interval 30 day))

-- e
select *, (PREMIUM - (PREMIUM * DISCOUNT / 100)) as PREMIUM_AFTER_DISCOUNT
from T_POLICY
where (PREMIUM - (PREMIUM * DISCOUNT / 100)) < 1000000
order by PREMIUM_AFTER_DISCOUNT asc
