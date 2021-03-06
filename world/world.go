package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/Nightgunner5/dfgen"
	"io"
	"os"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	handle(err)
	defer f.Close()

	r, err := dfgen.NewReader(bufio.NewReaderSize(f, 0x10000))
	handle(err)

	s, err := r.ReadShort()
	handle(err)
	fmt.Println("Unk1:", s)

	n, err := r.ReadLong()
	handle(err)
	fmt.Println("Unk2:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk3:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk4:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk5:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk6:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk7:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk8:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk9:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk10:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk11:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk12:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk13:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk14:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk15:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk16:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk17:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk18:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk19:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk20:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk21:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk22:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk23:", int32(n))

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk24:", int32(n))

	has_name, err := r.ReadByte()
	handle(err)
	if has_name > 1 {
		panic("has_name ∉ {0, 1}")
	}
	if has_name == 1 {
		NameStruct(r)
	}

	b, err := r.ReadByte()
	handle(err)
	fmt.Println("Unk25:", b)

	s, err = r.ReadShort()
	handle(err)
	fmt.Println("Unk26:", s)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk27:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk28:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk29:", n)

	str, err := r.ReadString()
	handle(err)
	fmt.Println("World name:", str)

	for _, kind := range []string{"MATERIAL", "ITEM", "CREATURE", "INTERACTION"} {
		listListLength, err := r.ReadLong()
		handle(err)
		fmt.Println(kind, "raws:", listListLength)

		for i := 0; i < int(listListLength); i++ {
			listLength, err := r.ReadLong()
			handle(err)

			for j := 0; j < int(listLength); j++ {
				str, err = r.ReadString()
				handle(err)

				fmt.Println(i, j, str)
			}
		}
	}

	var stInorganic, stPlant, stBody, stBodyGloss, stCreature, stItem, stBuilding, stEntity, stWord, stSymbol, stTranslation, stColor, stShape, stColorPattern, stReaction, stMaterialTemplate, stTissueTemplate, stBodyDetailPlan, stCreatureVariation, stInteraction []string

	for _, t := range []struct {
		name string
		list *[]string
	}{
		{"INORGANIC", &stInorganic},
		{"PLANT", &stPlant},
		{"BODY", &stBody},
		{"BODYGLOSS", &stBodyGloss},
		{"CREATURE", &stCreature},
		{"ITEM", &stItem},
		{"BUILDING", &stBuilding},
		{"ENTITY", &stEntity},
		{"WORD", &stWord},
		{"SYMBOL", &stSymbol},
		{"TRANSLATION", &stTranslation},
		{"COLOR", &stColor},
		{"SHAPE", &stShape},
		{"COLOR_PATTERN", &stColorPattern},
		{"REACTION", &stReaction},
		{"MATERIAL_TEMPLATE", &stMaterialTemplate},
		{"TISSUE_TEMPLATE", &stTissueTemplate},
		{"BODY_DETAIL_PLAN", &stBodyDetailPlan},
		{"CREATURE_VARIATION", &stCreatureVariation},
		{"INTERACTION", &stInteraction},
	} {
		listLength, err := r.ReadLong()
		handle(err)
		fmt.Println(t.name, "string table:", listLength)

		*t.list = make([]string, listLength)

		for i := 0; i < int(listLength); i++ {
			str, err = r.ReadString()
			handle(err)

			(*t.list)[i] = str

			fmt.Println(i, str)
		}
	}

	listLength, err := r.ReadLong()
	handle(err)
	fmt.Println("Unk30:", listLength, "records")
	for i := 0; i < int(listLength); i++ {
		j, err := r.ReadLong()
		handle(err)

		k, err := r.ReadLong()
		handle(err)

		fmt.Println(i, j, k)
	}

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk31:", n)

	listLength, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk32:", listLength, "records")
	unk32Length := int(listLength)
	for i := 0; i < int(listLength); i++ {
		n, err = r.ReadLong()
		fmt.Println(i, n)
	}

	listLength, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk33:", listLength, "records")
	for i := 0; i < int(listLength); i++ {
		n, err = r.ReadLong()
		fmt.Println(i, n)
	}

	listLength, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk34:", listLength, "records")
	for i := 0; i < int(listLength); i++ {
		n, err = r.ReadLong()
		fmt.Println(i, n)
	}

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk35:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk36:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk37:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk38:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk39:", n)

	listLength, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk40:", listLength, "records")
	for i := 0; i < int(listLength); i++ {
		n, err = r.ReadLong()
		fmt.Println(i, n)
	}

	listLength, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk41:", listLength, "records")
	for i := 0; i < int(listLength); i++ {
		n, err = r.ReadLong()
		fmt.Println(i, n)
	}
	unk41Length := int(listLength)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk42:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk43:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk44:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk45:", n)

	n, err = r.ReadLong()
	handle(err)
	fmt.Println("Unk46:", n)

	for i := 0; i < unk41Length; i++ {
		for j := 0; j < 3; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47a:", j, int16(s))
		}

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk47b:", s)

		listLength, err = r.ReadLong()
		handle(err)
		fmt.Println(i, "Unk47c:", listLength, "records")
		for j := 0; j < int(listLength); j++ {
			for k := 0; k < 3; k++ {
				s, err = r.ReadShort()
				handle(err)
				fmt.Println(i, "Unk47c:", j, k, s)
			}
		}

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk47d:", s)

		for j := 0; j < 3; j++ {
			n, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk47e:", j, int32(n))
		}

		for j := 0; j < 3; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47f:", j, int16(s))
		}

		for j := 0; j < 3; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47g:", j, int16(s))
		}

		b, err = r.ReadByte()
		handle(err)
		fmt.Println(i, "Unk47h:", b)

		for j := 0; j < 9; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47i:", j, int16(s))
		}

		for j := 0; j < 3; j++ {
			n, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk47j:", j, int32(n))
		}

		for j := 0; j < 3; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47k:", j, int16(s))
		}

		for j := 0; j < 3; j++ {
			n, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk47l:", j, int32(n))
		}

		for j := 0; j < 11; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk47m:", j, int16(s))
		}

		str, err = r.ReadString()
		handle(err)
		fmt.Println(i, "Unk47n:", str)
	}

	for i := 0; i < unk32Length; i++ {
		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk48a:", s)

		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, "Unk48b:", n)

		str, err = r.ReadString()
		handle(err)
		fmt.Println(i, "Unk48c:", str)

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk48d:", s)

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk48e:", s)

		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, "Unk48f:", n)

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Unk48g:", s)

		has_name, err := r.ReadByte()
		handle(err)
		if has_name > 1 {
			panic("has_name ∉ {0, 1}")
		}
		if has_name == 1 {
			NameStruct(r)
		}

		s, err = r.ReadShort()
		handle(err)
		fmt.Println(i, "Creature type:", s)
		fmt.Println("(", stCreature[s], ")")

		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, "Unk48h:", n)

		for j := 0; j < 19; j++ {
			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48i", j, "a:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				s, err = r.ReadShort()
				handle(err)
				fmt.Println(i, "Unk48i", j, "a:", k, s)
			}

			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48i", j, "b:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				n, err = r.ReadLong()
				handle(err)
				fmt.Println(i, "Unk48i", j, "b:", k, n)
			}
		}

		for j := 0; j < 12; j++ {
			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48j", j, "a:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				n, err = r.ReadLong()
				handle(err)
				fmt.Println(i, "Unk48j", j, "a:", k, n)
			}

			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48j", j, "b:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				s, err = r.ReadShort()
				handle(err)
				fmt.Println(i, "Unk48j", j, "b:", k, s)
			}
		}

		for j := 0; j < 2; j++ {
			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48k", j, "a:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				s, err = r.ReadShort()
				handle(err)
				fmt.Println(i, "Unk48k", j, "a:", k, s)
			}

			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48k", j, "b:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				n, err = r.ReadLong()
				handle(err)
				fmt.Println(i, "Unk48k", j, "b:", k, n)
			}
		}

		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, "Unk48l:", n)

		for j := 0; j < 9; j++ {
			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48m", j, "a:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				s, err = r.ReadShort()
				handle(err)
				fmt.Println(i, "Unk48m", j, "a:", k, s)
			}

			listLength, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48m", j, "b:", listLength, "records")
			for k := 0; k < int(listLength); k++ {
				n, err = r.ReadLong()
				handle(err)
				fmt.Println(i, "Unk48m", j, "b:", k, n)
			}
		}

		for j := 0; j < 3; j++ {
			s, err = r.ReadShort()
			handle(err)
			fmt.Println(i, "Unk48n", j, "a:", int16(s))

			n, err = r.ReadLong()
			handle(err)
			fmt.Println(i, "Unk48n", j, "b:", n)
		}
		break
	}

	x := hex.Dumper(os.Stdout)
	defer x.Close()

	_, err = io.Copy(x, r.R)
	handle(err)
}

func NameStruct(r dfgen.Reader) {
	fmt.Println("Name struct:")

	str, err := r.ReadString()
	handle(err)
	fmt.Println("Name:", str)

	str, err = r.ReadString()
	handle(err)
	fmt.Println("Nickname:", str)

	var (
		wordIndices [7]int32
		wordForms   [7]string
	)

	for i := range wordIndices {
		n, err := r.ReadLong()
		handle(err)
		wordIndices[i] = int32(n)
	}
	for i := range wordForms {
		s, err := r.ReadShort()
		handle(err)
		wordForms[i] = []string{"noun[singular]", "noun[plural]", "adjective", "prefix", "present[1st]", "present[3rd]", "preterite", "past[participle]", "present[participle]"}[s]
	}
	fmt.Println("Word indices:", wordIndices)
	fmt.Println("Word indices:", wordForms)

	n, err := r.ReadLong()
	handle(err)
	fmt.Println("Language:", int32(n))

	s, err := r.ReadShort()
	handle(err)
	fmt.Println("Unknown:", s)
}
