package client

import "os"

func PrepareVCS(vcs []vcsDataResponse) (map[string]table, error) {
	c, err := GetScriptedColumns()
	if err != nil {
		return nil, err
	}

	r := make(map[string]table, len(c))

	for k, v := range c {
		if _, ok := r[k]; !ok {
			r[k] = table{Column: append(v, "sys_id")}
		}
	}

	for _, v := range vcs {
		if entry, ok := r[v.TableName]; ok {
			entry.RecordID = append(entry.RecordID, v.RecordID)
			r[v.TableName] = entry
		}
	}

	return r, nil
}

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
