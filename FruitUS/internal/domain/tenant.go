package domain

type Tenant struct {
    ID   string
    Name string
    // Add other tenant-related fields as necessary
}

func (t *Tenant) GetID() string {
    return t.ID
}

func (t *Tenant) GetName() string {
    return t.Name
}

// Add other methods related to Tenant as necessary