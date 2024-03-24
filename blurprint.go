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

func (b *Blueprint) Varchar(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: VARCHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Char(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: CHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Text(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TEXT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Date(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Timestamp(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TIMESTAMP,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) DateTime(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATETIME,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Enum(name string, options []string, props *EnumColumnProps) {
	dataType := SQLTableProp{
		Type:        ENUM,
		EnumOptions: options,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Int(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: INT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Serial(name string) {
	dataType := SQLTableProp{
		Type: SERIAL,
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) BigSerial(name string) {
	dataType := SQLTableProp{
		Type:     BIGSERIAL,
		Unsigned: true,
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Tinyint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: TINYINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.Columns = append(b.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (b *Blueprint) Mediumint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: MEDIUMINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.Columns = append(b.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (b *Blueprint) Bigint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BIGINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.Columns = append(b.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (b *Blueprint) Boolean(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BOOL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Float(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: FLOAT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Double(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Real(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: REAL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) DoublePrecision(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE_PRECISION,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Increment(name string) {
	dataType := SQLTableProp{
		Type:       SERIAL,
		PrimaryKey: true,
	}

	if b.Dialect == MYSQL {
		dataType.Type = INT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) BigIncrement(name string) {
	dataType := SQLTableProp{
		Type:       BIGSERIAL,
		PrimaryKey: true,
	}

	if b.Dialect == MYSQL {
		dataType.Type = BIGINT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	b.AddColumn(name, dataType)
}

func (b *Blueprint) Uuid(name string, props *UUIDColumnProps) {
	dataType := SQLTableProp{
		Type:    UUID,
		Default: "gen_random_uuid()",
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	if b.Dialect == MYSQL {
		dataType.Type = VARCHAR
		dataType.Default = "uuid()"
		dataType.Size = 36
	}

	b.AddColumn(name, dataType)
}
