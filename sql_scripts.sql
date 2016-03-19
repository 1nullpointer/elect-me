CREATE TABLE `votes` (
  `ID` int(11) DEFAULT NULL,
  `Ward` int(11) DEFAULT NULL,
  `Division` int(11) DEFAULT NULL,
  `Type` char(1) DEFAULT NULL,
  `Vote` int(11) DEFAULT NULL,
  `OfficeID` int(11) DEFAULT NULL,
  `PartyID` int(11) DEFAULT NULL,
  `CandidateID` int(11) DEFAULT NULL,
  KEY `votes_OfficeID` (`OfficeID`),
  KEY `votes_PartyID` (`PartyID`),
  KEY `votes_CandidateID` (`CandidateID`)
);

CREATE TABLE `offices` (
  `OfficeID` int(11) NOT NULL DEFAULT '0',
  `Name` varchar(500) DEFAULT NULL,
  `OfficeTypeID` int(11) DEFAULT NULL,
  PRIMARY KEY (`OfficeID`)
);

CREATE TABLE `candidates` (
  `CandidateID` int(11) NOT NULL DEFAULT '0',
  `Name` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`CandidateID`)
);

CREATE TABLE `parties` (
  `PartyID` int(11) NOT NULL DEFAULT '0',
  `Name` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`PartyID`)
);

CREATE TABLE `OfficeTypes` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`ID`)
);

insert into OfficeTypes (Name) values ('City'), ('District'), ('Referendum');


update offices set OfficeTypeID = case when Name like '%DISTRICT%' then 2 
                                    when Name like '%QUESTION%' then 3 
                                    else 1 end;

select t.Office, max(t.VoteCount) as WinnerVoteCount
from
    (
    select o.Name as Office, c.CandidateID, sum(v.vote) as VoteCount
    from votes v
        left join candidates c on v.CandidateID = c.CandidateID
        left join offices o on v.OfficeID = o.OfficeID
        left join OfficeTypes ot on o.OfficeTypeID = ot.ID
    where ot.Name = 'City' and o.Name not like 'RETENTION%' and o.Name <> 'No Vote'
    group by o.OfficeID, c.CandidateID
    ) as t
group by t.Office
order by WinnerVoteCount
limit 10;
