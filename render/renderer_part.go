package render

import (
	"blockexchange/types"
	"sort"

	"github.com/fogleman/gg"
)

type Block struct {
	X     int
	Y     int
	Z     int
	Color *Color
	Order int
}

type PartRenderer struct {
	Schemapart          *types.SchemaPart
	Mapblock            *ParsedSchemaPart
	Colormapping        map[string]*Color
	NodeIDStringMapping map[int]string
	Blocks              []*Block
	MaxX                int
	MaxY                int
	MaxZ                int
	YMultiplier         int
	XMultiplier         int
	Size                float64
	OffsetX             float64
	OffsetY             float64
}

func NewPartRenderer(schemapart *types.SchemaPart, mapblock *ParsedSchemaPart, cm map[string]*Color, size, offset_x, offset_y float64) *PartRenderer {
	// reverse index
	idm := make(map[int]string)
	for k, v := range mapblock.Meta.NodeMapping {
		idm[v] = k
	}
	return &PartRenderer{
		Schemapart:          schemapart,
		Mapblock:            mapblock,
		Blocks:              make([]*Block, 0),
		Colormapping:        cm,
		NodeIDStringMapping: idm,
		MaxX:                mapblock.Meta.Size.X - 1,
		MaxY:                mapblock.Meta.Size.Y - 1,
		MaxZ:                mapblock.Meta.Size.Z - 1,
		YMultiplier:         mapblock.Meta.Size.Z,
		XMultiplier:         mapblock.Meta.Size.Y * mapblock.Meta.Size.Z,
		Size:                size,
		OffsetX:             offset_x,
		OffsetY:             offset_y,
	}
}

func (r *PartRenderer) GetImagePos(x, y, z float64) (float64, float64) {
	xpos := r.OffsetX + (r.Size * x) - (r.Size * z)
	ypos := r.OffsetY - (r.Size * tan30 * x) - (r.Size * tan30 * z) - (r.Size * sqrt3div2 * y)

	return xpos, ypos
}

func (r *PartRenderer) GetColorAtPos(x, y, z int) *Color {
	if x > r.MaxX || y > r.MaxY || z > r.MaxZ || x < 0 || y < 0 || z < 0 {
		return nil
	}

	index := z + (y * r.YMultiplier) + (x * r.XMultiplier)
	nodeid := int(r.Mapblock.NodeIDS[index])
	nodename := r.NodeIDStringMapping[nodeid]
	color := r.Colormapping[nodename]
	return color
}

func (r *PartRenderer) ProbePosition(x, y, z int) {
	color := r.GetColorAtPos(x, y, z)
	if color != nil {
		block := Block{
			X:     x,
			Y:     y,
			Z:     z,
			Color: color,
			Order: y + ((r.MaxX - x) * r.MaxX) + ((r.MaxZ - z) + r.MaxZ),
		}

		r.Blocks = append(r.Blocks, &block)
		return
	}

	next_x := x + 1
	next_y := y - 1
	next_z := z + 1

	if next_x > r.MaxX || next_z > r.MaxZ || next_y < 0 {
		// mapblock ends
		return
	}

	r.ProbePosition(next_x, next_y, next_z)
}

func (r *PartRenderer) DrawBlock(dc *gg.Context, block *Block) {
	x, y := r.GetImagePos(float64(block.X), float64(block.Y), float64(block.Z))
	radius := r.Size

	// right side
	dc.MoveTo(radius+x, (radius*tan30)+y)
	dc.LineTo(x, (radius*sqrt3div2)+y)
	dc.LineTo(x, y)
	dc.LineTo(radius+x, -(radius*tan30)+y)
	dc.ClosePath()
	dc.SetRGB255(block.Color.Red, block.Color.Green, block.Color.Blue)
	dc.Fill()

	// left side
	dc.MoveTo(x, (radius*sqrt3div2)+y)
	dc.LineTo(-radius+x, (radius*tan30)+y)
	dc.LineTo(-radius+x, -(radius*tan30)+y)
	dc.LineTo(x, y)
	dc.ClosePath()
	AdjustAndFill(dc, block.Color.Red, block.Color.Green, block.Color.Blue, -20)
	dc.Fill()

	// top side
	dc.MoveTo(-radius+x, -(radius*tan30)+y)
	dc.LineTo(x, -(radius*sqrt3div2)+y)
	dc.LineTo(radius+x, -(radius*tan30)+y)
	dc.LineTo(x, y)
	dc.ClosePath()
	AdjustAndFill(dc, block.Color.Red, block.Color.Green, block.Color.Blue, 20)
	dc.Fill()
}

func (r *PartRenderer) RenderSchemaPart(dc *gg.Context) error {

	for y := 0; y < r.MaxY; y++ {
		// right side
		for x := r.MaxX; x >= 1; x-- {
			r.ProbePosition(x, y, 0)
		}

		// left side
		for z := r.MaxZ; z >= 0; z-- {
			r.ProbePosition(0, y, z)
		}
	}

	// top side
	for z := r.MaxZ; z >= 0; z-- {
		for x := r.MaxX; x >= 0; x-- {
			r.ProbePosition(x, r.MaxY, z)
		}
	}

	sort.Slice(r.Blocks, func(i int, j int) bool {
		return r.Blocks[i].Order < r.Blocks[j].Order
	})

	for _, block := range r.Blocks {
		r.DrawBlock(dc, block)
	}

	return nil
}
