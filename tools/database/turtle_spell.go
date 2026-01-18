package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/isfir/wowsims-turtle/sim/core/proto"
)

func ParseTurtleSpellDB(spellCSV, spellIconCSV string) []*proto.IconData {
	spellIconMap := parseSpellIconCSV(spellIconCSV)
	return parseSpellCSV(spellCSV, spellIconMap)
}

func parseSpellIconCSV(csvData string) map[int32]string {
	r := csv.NewReader(strings.NewReader(csvData))
	// Skip header
	if _, err := r.Read(); err != nil {
		log.Fatalf("Cannot read spell icon csv header: %v", err)
	}

	iconMap := make(map[int32]string)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cannot read spell icon csv row: %v", err)
		}
		if len(row) < 2 {
			continue
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Printf("Invalid spell icon id %s: %v", row[0], err)
			continue
		}
		texture := row[1]
		// Extract texture name from path like "Interface\Icons\Spell_Shadow_BlackPlague"
		parts := strings.Split(texture, "\\")
		if len(parts) > 0 {
			texture = parts[len(parts)-1]
		}
		texture = strings.ToLower(texture)
		iconMap[int32(id)] = texture
	}
	return iconMap
}

func parseSpellCSV(csvData string, iconMap map[int32]string) []*proto.IconData {
	r := csv.NewReader(strings.NewReader(csvData))
	headers, err := r.Read()
	if err != nil {
		log.Fatalf("Cannot read spell csv header: %v", err)
	}

	// Map column names to indices
	colIdx := make(map[string]int)
	for i, name := range headers {
		colIdx[name] = i
	}

	requiredCols := []string{"id", "spellIconId", "name_enUS", "subtext_enUS", "effectApplyAura_1", "effectApplyAura_2", "effectApplyAura_3"}
	for _, col := range requiredCols {
		if _, ok := colIdx[col]; !ok {
			log.Fatalf("Missing required column %s in spell csv", col)
		}
	}

	var spells []*proto.IconData
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cannot read spell csv row: %v", err)
		}

		id, err := strconv.Atoi(row[colIdx["id"]])
		if err != nil {
			log.Printf("Invalid spell id %s: %v", row[colIdx["id"]], err)
			continue
		}

		spellIconId, err := strconv.Atoi(row[colIdx["spellIconId"]])
		if err != nil {
			log.Printf("Invalid spellIconId %s for spell %d: %v", row[colIdx["spellIconId"]], id, err)
			continue
		}

		name := row[colIdx["name_enUS"]]
		if name == "" {
			continue
		}

		icon := iconMap[int32(spellIconId)]
		if icon == "" {
			// Fallback to activeIconId if column exists
			if activeIdx, ok := colIdx["activeIconId"]; ok {
				if activeIdStr := row[activeIdx]; activeIdStr != "" {
					if activeId, err := strconv.Atoi(activeIdStr); err == nil {
						icon = iconMap[int32(activeId)]
					}
				}
			}
		}

		// Determine rank from subtext_enUS (e.g., "Rank 2")
		rank := int32(0)
		subtext := strings.TrimSpace(row[colIdx["subtext_enUS"]])
		if strings.HasPrefix(subtext, "Rank ") {
			rankStr := strings.TrimPrefix(subtext, "Rank ")
			if rankNum, err := strconv.Atoi(rankStr); err == nil {
				rank = int32(rankNum)
			}
		}

		// Determine if spell has buff (any non-zero effectApplyAura)
		hasBuff := false
		for i := 1; i <= 3; i++ {
			colName := fmt.Sprintf("effectApplyAura_%d", i)
			valStr := row[colIdx[colName]]
			if valStr != "" && valStr != "0" {
				hasBuff = true
				break
			}
		}

		spells = append(spells, &proto.IconData{
			Id:      int32(id),
			Name:    name,
			Icon:    icon,
			Rank:    rank,
			HasBuff: hasBuff,
		})
	}
	return spells
}
