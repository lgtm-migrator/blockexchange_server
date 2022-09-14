package worldedit

import (
	"errors"

	lua "github.com/yuin/gopher-lua"
)

func Parse(data []byte) ([]*WEEntry, error) {
	if data[0] != byte('5') || data[1] != byte(':') {
		return nil, errors.New("invalid format")
	}

	L := lua.NewState()
	defer L.Close()

	fn, err := L.LoadString(string(data[2:]))
	if err != nil {
		return nil, err
	}

	err = L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		return nil, err
	}

	lv := L.Get(-1)
	tbl, ok := lv.(*lua.LTable)
	if !ok {
		return nil, errors.New("root table not found")
	}

	entry_count := L.ObjLen(tbl)
	entries := make([]*WEEntry, entry_count)

	for i := 1; i <= entry_count; i++ {
		entrytbl := L.GetTable(tbl, lua.LNumber(i))

		// common entries
		entry, err := parseWEEntry(L, entrytbl)
		if err != nil {
			return nil, err
		}
		entries[i-1] = entry

		metatbl := L.GetTable(entrytbl, lua.LString("meta"))
		if metatbl != lua.LNil {
			entry.Meta = &WEMeta{
				Fields:    make(WEFields),
				Inventory: make(WEInventory),
			}

			// fields
			fieldstbl := L.GetTable(metatbl, lua.LString("fields")).(*lua.LTable)
			if fieldstbl != lua.LNil {
				fieldstbl.ForEach(func(key, value lua.LValue) {
					entry.Meta.Fields[key.String()] = value.String()
				})
			}

			invtbl := L.GetTable(metatbl, lua.LString("inventory")).(*lua.LTable)
			if invtbl != lua.LNil {
				invtbl.ForEach(func(key, value lua.LValue) {
					stacks := make([]string, 0)
					stacktbl := value.(*lua.LTable)
					stacktbl.ForEach(func(_, stack lua.LValue) {
						stacks = append(stacks, stack.String())
					})
					entry.Meta.Inventory[key.String()] = stacks
				})
			}
		}

	}

	return entries, nil
}

func parseWEEntry(L *lua.LState, tbl lua.LValue) (*WEEntry, error) {
	e := &WEEntry{}

	name, ok := L.GetTable(tbl, lua.LString("name")).(lua.LString)
	if ok {
		e.Name = name.String()
	}

	param2, ok := L.GetTable(tbl, lua.LString("param2")).(lua.LNumber)
	if ok {
		e.Param2 = byte(param2)
	}

	param1, ok := L.GetTable(tbl, lua.LString("param1")).(lua.LNumber)
	if ok {
		e.Param1 = byte(param1)
	}

	x, ok := L.GetTable(tbl, lua.LString("x")).(lua.LNumber)
	if ok {
		e.X = int(x)
	}

	y, ok := L.GetTable(tbl, lua.LString("y")).(lua.LNumber)
	if ok {
		e.Y = int(y)
	}

	z, ok := L.GetTable(tbl, lua.LString("z")).(lua.LNumber)
	if ok {
		e.Z = int(z)
	}

	return e, nil
}