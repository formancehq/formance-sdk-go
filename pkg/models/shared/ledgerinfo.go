// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Storage struct {
	Migrations []MigrationInfo `json:"migrations,omitempty"`
}

func (o *Storage) GetMigrations() []MigrationInfo {
	if o == nil {
		return nil
	}
	return o.Migrations
}

type LedgerInfo struct {
	Name    *string  `json:"name,omitempty"`
	Storage *Storage `json:"storage,omitempty"`
}

func (o *LedgerInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LedgerInfo) GetStorage() *Storage {
	if o == nil {
		return nil
	}
	return o.Storage
}
