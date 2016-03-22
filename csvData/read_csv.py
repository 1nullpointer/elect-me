import pandas as pd

df_2015 = pd.DataFrame.from_csv('2015_general.csv', header=None, index_col=False)

df_2015.columns = ['Ward', 'Division', 'Type', 'Office', 'Candidate', 'Party', 'Vote']

#del df_2015['7']
#del df_2015['8']

office_df = df_2015['Office'].drop_duplicates()
candidate_df = df_2015['Candidate'].drop_duplicates()
party_df = df_2015['Party'].drop_duplicates()

inverse = lambda d: {v: k for k, v in d.items()}

office = inverse(office_df.to_dict())
candidate = inverse(candidate_df.to_dict())
party = inverse(party_df.to_dict())


df_2015['OfficeID'] = df_2015['Office'].apply(lambda x: office[x])
df_2015['PartyID'] = df_2015['Party'].apply(lambda x: party[x])
df_2015['CandidateID'] = df_2015['Candidate'].apply(lambda x: candidate[x])

del df_2015['Office']
del df_2015['Candidate']
del df_2015['Party']

df_2015.to_csv('votes2015.csv')
office_df.to_csv('office2015.csv')
party_df.to_csv('party2015.csv')
candidate_df.to_csv('candidate2015.csv')