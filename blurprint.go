package gomigrator

type Blueprint struct {
	Columns []TableColumn
	Dialect SQLDialect
}

func (b *Blueprint) AddColumn(name string, props SQLTableProp) {
	b.Columns = append(b.Columns, TableColumn{
		Name:     name,
		Property: &props,
	})
}

func (t *Blueprint) Varchar(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: VARCHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Char(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: CHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Text(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TEXT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Date(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Timestamp(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TIMESTAMP,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) DateTime(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATETIME,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Enum(name string, options []string, props *EnumColumnProps) {
	dataType := SQLTableProp{
		Type:        ENUM,
		EnumOptions: options,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Int(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: INT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Serial(name string) {
	dataType := SQLTableProp{
		Type: SERIAL,
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) BigSerial(name string) {
	dataType := SQLTableProp{
		Type:     BIGSERIAL,
		Unsigned: true,
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Tinyint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: TINYINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Blueprint) Mediumint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: MEDIUMINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Blueprint) Bigint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BIGINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Blueprint) Boolean(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BOOL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Float(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: FLOAT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Double(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Real(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: REAL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) DoublePrecision(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE_PRECISION,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Increment(name string) {
	dataType := SQLTableProp{
		Type:       SERIAL,
		PrimaryKey: true,
	}

	if t.Dialect == MYSQL {
		dataType.Type = INT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) BigIncrement(name string) {
	dataType := SQLTableProp{
		Type:       BIGSERIAL,
		PrimaryKey: true,
	}

	if t.Dialect == MYSQL {
		dataType.Type = BIGINT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	t.AddColumn(name, dataType)
}

func (t *Blueprint) Uuid(name string, props *UUIDColumnProps) {
	dataType := SQLTableProp{
		Type:    UUID,
		Default: "gen_random_uuid()",
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	if t.Dialect == MYSQL {
		dataType.Type = VARCHAR
		dataType.Default = "uuid()"
		dataType.Size = 36
	}

	t.AddColumn(name, dataType)
}
