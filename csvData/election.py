import pandas as pd
from csv import Sniffer

class Election:
    
    def __init__(self, year, type, scope):
        self.year = year
        self.type = type
        self.scope = scope
        self.data = None
        
    def get_data(self, path_to_csv, column_names=None):
        column_names = column_names if column_names else 'infer'
        df = pd.DataFrame.from_csv(path_to_csv, index_col=False, header=column_names)
        if self._csv_has_index(df):
            df = pd.DataFrame.from_csv(path_to_csv, index_col=True, header=column_names)
        self.data = df
        self._parse_info(path_to_csv)
    
    def _parse_info(self, path_to_csv):
        info = path_to_csv.split('_')
        self.year = info[0]
        self.type = info[1]
        self.scope = info[3]
    
    def add_data(self, other):
        assert type(self) and type(other)
        assert not self.data.empty and not other.data.empty
        assert self.year == other.year
        assert self.type == other.type
        assert self.scope == other.scope
        
        self.data = pd.concat([self.data, other.data])
    
    def get_lowest_votes_positions(self, start=0, end=10):
        if self._columns_are_valid():
                grouped = self.data.groupby(['Position'])['Votes'].max()
                grouped.sort()
                return grouped[start:end]
        else:
            raise Exception('Rename columns before calculating votes')
            
    
    def rename_columns(self, renaming_dict={'Category': 'Position', 
                                            'Selection': 'Candidate'}):
        assert not self.data.empty
        self.data.rename(columns=renaming_dict, inplace=True)
    
    def _csv_has_index(self, df):
        return len(df[df.columns[0]]) == len(df[df.columns[0]].unique())    
    
    def _columns_are_valid(self):
        for column in ['Position', 'Candidate', 'Votes']:
            if column not in self.data.columns:
                return False
        return True