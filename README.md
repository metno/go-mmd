# What is go-mmd?

go-mmd is a go-library for handling the MMD metadata format: https://github.com/steingod/mmd

Separate structs for decoding and encoding, due to go issue with namespaces: https://github.com/golang/go/issues/9519

## Decode MMD

```go
var mmdAromeArcticXMLDocument = `
<mmd:mmd xmlns:mmd="http://www.met.no/schema/mmd" xmlns:gml="http://www.opengis.net/gml">
  <mmd:metadata_identifier>no.met:72170bc4-c0e9-4c57-adc9-2fecb1ddfa84</mmd:metadata_identifier>
  <mmd:title xml:lang="en">Arome-Arctic 2.5Km deterministic 2023-08-10T06:00:00Z + 66 hours</mmd:title>
  <mmd:abstract xml:lang="en">This file contains output from Control member. Contains most forecast parameters on surface, model level and pressure levels The Norwegian Meteorological Institute has provided forecasts of the weather in the High North for a long time. Our Development Centre for Weather Forecasting has developed a high-resolution model that forecasts the weather in this area for the next three days. The model is called AROME-Arctic and has been running since November 15th, 2015. The geographic resolution is 2.5 kilometres, which is the same as in the model that the Norwegian Meteorological Institute uses to forecast the weather elsewhere in Norway. AROME-Arctic gives a detailed representation of processes on the ground and in the atmosphere. In addition to Svalbard and northern Norway, the model also covers a large area of sea, which is a challenge due to the sparsity of observations. Thus, satellite observations are particularly important for this model. </mmd:abstract>
  <mmd:abstract xml:lang="no">Denne filen inneholder utdata fra kontrollmedlem. Inneholder de fleste prognoseparametere på overflate, modellnivå og trykknivåer Meteorologisk institutt har lenge levert varsler om været i nordområdene. Vårt utviklingssenter for værvarsling har utviklet en høyoppløselig modell som varsler været i dette området for de neste tre dagene. Modellen heter AROME-Arctic og har vært i drift siden 15. november 2015. Den geografiske oppløsningen er 2,5 kilometer, som er den samme som i modellen Meteorologisk Institutt bruker for å varsle været andre steder i Norge. AROME-Arctic gir en detaljert fremstilling av prosesser på bakken og i atmosfæren. I tillegg til Svalbard og Nord-Norge dekker modellen også et stort havområde, noe som er en utfordring på grunn begrenset mengde in-situ overflateobservasjoner. Derfor er satellittobservasjoner spesielt viktige for denne modellen. </mmd:abstract>
  <mmd:metadata_status>Active</mmd:metadata_status>
  <mmd:dataset_production_status>Complete</mmd:dataset_production_status>
  <mmd:collection>METNCS</mmd:collection>
</mmd:mmd>
`

var metadata DecodeMetadataMMD
err := xml.Unmarshal([]byte(mmdAromeArcticXMLDocument), &metadata)
```

### Encode MMD

```go
var MMDExample = MetadataMMD{
    Title: &MMDTitle{
        Title:    "Arome Arctic",
        Language: "en",
    },
}

encoded, err := xml.Marshal(MMDExample)
```
