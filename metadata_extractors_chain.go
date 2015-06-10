package db

import "reflect"

type MetadataExtractorChain []MetadataExtractor

func NewKeyResolverExtractorChain(metadata *Metadata) MetadataExtractorChain {
	return MetadataExtractorChain{
		IntIdExtractor{metadata},
		StringIdExtractor{metadata},
		KindExtractor{metadata},
		HasParentExtractor{metadata},
	}
}

// TODO Move related Tests from key_resolver_tests to metadata_extractors_chain_test
func (this MetadataExtractorChain) ExtractFrom(e Entity) error {
	elemType := reflect.TypeOf(e).Elem()
	elemValue := reflect.ValueOf(e).Elem()

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		value := elemValue.Field(i)
		for _, extractor := range this {
			if extractor.Accept(field) {
				if err := extractor.Extract(e, field, value); err != nil {
					return err
				}
			}
		}
	}

	return nil
}