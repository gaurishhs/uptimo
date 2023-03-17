package models

type StatusPage struct {
  ID        uuid.UUID `db:"id"`
  Name      string    `db:"name"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
  Monitors  []Monitor `has_many:"monitors"`

}
