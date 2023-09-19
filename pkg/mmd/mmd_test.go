package mmd

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMMDDecoding(t *testing.T) {
	var metadata DecodeMetadataMMD
	err := xml.Unmarshal([]byte(mmdAromeArcticEncoded), &metadata)
	require.Nil(t, err)

	require.Equal(t, "Arome-Arctic 2.5Km deterministic 2023-08-10T06:00:00Z + 66 hours", metadata.Title.Title)
	require.Equal(t, "2023-08-10T06:00:00Z", metadata.TemporalExtent.StartDate)

	require.Equal(t, 1, len(metadata.ISOTopicCategories), "ISOTopicCatagories should list one value")
	require.Equal(t, "climatologyMeteorologyAtmosphere", metadata.ISOTopicCategories[0])

	require.Equal(t, 47006.87, metadata.StorageInformation.FileSize)

	expectedKeyword := "Weather and climate"
	var keywordFound bool
	for _, v := range metadata.Keywords {
		for _, k := range v.Keyword {
			if k == expectedKeyword {
				keywordFound = true
			}
		}
	}
	require.True(t, keywordFound)
}

func TestMMDEncoding(t *testing.T) {
	encoded, err := xml.Marshal(MMDExample)
	require.Nil(t, err)

	require.Contains(t, string(encoded), fmt.Sprintf("<mmd:title xml:lang=%q>Arome Arctic", "en"))
	require.Contains(t, string(encoded),
		"<mmd:data_center><mmd:data_center_name><mmd:short_name>MET Norway</mmd:short_name><mmd:long_name>Norwegian Meteorological Institute</mmd:long_name></mmd:data_center_name><mmd:data_center_url>https://data.met.no</mmd:data_center_url></mmd:data_center>")
}

var mmdAromeArcticEncoded = `
<?xml version="1.0"?>
<mmd:mmd xmlns:mmd="http://www.met.no/schema/mmd" xmlns:gml="http://www.opengis.net/gml">
  <mmd:metadata_identifier>no.met:72170bc4-c0e9-4c57-adc9-2fecb1ddfa84</mmd:metadata_identifier>
  <mmd:title xml:lang="en">Arome-Arctic 2.5Km deterministic 2023-08-10T06:00:00Z + 66 hours</mmd:title>
  <mmd:abstract xml:lang="en">This file contains output from Control member. Contains most forecast parameters on surface, model level and pressure levels The Norwegian Meteorological Institute has provided forecasts of the weather in the High North for a long time. Our Development Centre for Weather Forecasting has developed a high-resolution model that forecasts the weather in this area for the next three days. The model is called AROME-Arctic and has been running since November 15th, 2015. The geographic resolution is 2.5 kilometres, which is the same as in the model that the Norwegian Meteorological Institute uses to forecast the weather elsewhere in Norway. AROME-Arctic gives a detailed representation of processes on the ground and in the atmosphere. In addition to Svalbard and northern Norway, the model also covers a large area of sea, which is a challenge due to the sparsity of observations. Thus, satellite observations are particularly important for this model. </mmd:abstract>
  <mmd:abstract xml:lang="no">Denne filen inneholder utdata fra kontrollmedlem. Inneholder de fleste prognoseparametere på overflate, modellnivå og trykknivåer Meteorologisk institutt har lenge levert varsler om været i nordområdene. Vårt utviklingssenter for værvarsling har utviklet en høyoppløselig modell som varsler været i dette området for de neste tre dagene. Modellen heter AROME-Arctic og har vært i drift siden 15. november 2015. Den geografiske oppløsningen er 2,5 kilometer, som er den samme som i modellen Meteorologisk Institutt bruker for å varsle været andre steder i Norge. AROME-Arctic gir en detaljert fremstilling av prosesser på bakken og i atmosfæren. I tillegg til Svalbard og Nord-Norge dekker modellen også et stort havområde, noe som er en utfordring på grunn begrenset mengde in-situ overflateobservasjoner. Derfor er satellittobservasjoner spesielt viktige for denne modellen. </mmd:abstract>
  <mmd:metadata_status>Active</mmd:metadata_status>
  <mmd:dataset_production_status>Complete</mmd:dataset_production_status>
  <mmd:collection>METNCS</mmd:collection>
  <mmd:last_metadata_update>
    <mmd:update>
      <mmd:datetime>2023-08-10T09:54:58Z</mmd:datetime>
      <mmd:type>Created</mmd:type>
    </mmd:update>
  </mmd:last_metadata_update>
  <mmd:temporal_extent>
    <mmd:start_date>2023-08-10T06:00:00Z</mmd:start_date>
    <mmd:end_date>2023-08-13T00:00:00Z</mmd:end_date>
  </mmd:temporal_extent>
  <mmd:iso_topic_category>climatologyMeteorologyAtmosphere</mmd:iso_topic_category>
  <mmd:keywords vocabulary="GCMDSK">
    <mmd:keyword>WEATHER FORECAST</mmd:keyword>
    <mmd:resource>https://gcmd.earthdata.nasa.gov/kms/concepts/concept_scheme/sciencekeywords</mmd:resource>
    <mmd:separator></mmd:separator>
  </mmd:keywords>
  <mmd:keywords vocabulary="GEMET">
    <mmd:keyword>Meteorological geographical features</mmd:keyword>
    <mmd:keyword>Atmospheric conditions</mmd:keyword>
    <mmd:resource>https://inspire.ec.europa.eu/theme</mmd:resource>
    <mmd:separator></mmd:separator>
  </mmd:keywords>
  <mmd:keywords vocabulary="NORTHEMES">
    <mmd:keyword>Weather and climate</mmd:keyword>
    <mmd:resource>https://register.geonorge.no/metadata-kodelister/nasjonal-temainndeling</mmd:resource>
    <mmd:separator></mmd:separator>
  </mmd:keywords>
  <mmd:geographic_extent>
    <mmd:rectangle srsName="EPSG:4326">
      <mmd:north>88.0</mmd:north>
      <mmd:south>62.0</mmd:south>
      <mmd:east>80.0</mmd:east>
      <mmd:west>-18.0</mmd:west>
    </mmd:rectangle>
  </mmd:geographic_extent>
  <mmd:dataset_language>en</mmd:dataset_language>
  <mmd:operational_status>Not available</mmd:operational_status>
  <mmd:use_constraint>
    <mmd:identifier>CC-BY-4.0</mmd:identifier>
    <mmd:resource>https://spdx.org/licenses/CC-BY-4.0</mmd:resource>
  </mmd:use_constraint>
  <mmd:personnel>
    <mmd:role>Investigator</mmd:role>
    <mmd:name>Senter for utvikling av værvarslingstjenesten</mmd:name>
    <mmd:email>suv-arctic@met.no</mmd:email>
    <mmd:organisation>Norwegian Meteorological Institute</mmd:organisation>
  </mmd:personnel>
  <mmd:data_center>
    <mmd:data_center_name>
      <mmd:short_name>MET Norway</mmd:short_name>
      <mmd:long_name>Norwegian Meteorological Institute</mmd:long_name>
    </mmd:data_center_name>
    <mmd:data_center_url>https://data.met.no</mmd:data_center_url>
  </mmd:data_center>
  <mmd:data_access>
    <mmd:type>OPeNDAP</mmd:type>
    <mmd:description>Open-source Project for a Network Data Access Protocol</mmd:description>
    <mmd:resource>https://thredds.met.no/thredds/dodsC/aromearcticarchive/2023/08/10/arome_arctic_det_2_5km_20230810T06Z.nc</mmd:resource>
  </mmd:data_access>
  <mmd:data_access>
    <mmd:type>HTTP</mmd:type>
    <mmd:description>Direct download of file</mmd:description>
    <mmd:resource>https://thredds.met.no/thredds/fileServer/aromearcticarchive/2023/08/10/arome_arctic_det_2_5km_20230810T06Z.nc</mmd:resource>
  </mmd:data_access>
  <mmd:storage_information>
    <mmd:file_name>arome_arctic_det_2_5km_20230810T06Z.nc</mmd:file_name>
    <mmd:file_location>/lustre/storeA/project/metproduction/work/run06/det/2023/08/10</mmd:file_location>
    <mmd:file_format>NetCDF-CF</mmd:file_format>
    <mmd:file_size unit="MB">47006.87</mmd:file_size>
  </mmd:storage_information>
  <mmd:related_information>
    <mmd:type>Project home page</mmd:type>
    <mmd:description></mmd:description>
    <mmd:resource>https://www.met.no/en/projects/The-weather-model-AROME-Arctic</mmd:resource>
  </mmd:related_information>
  <mmd:spatial_representation>grid</mmd:spatial_representation>
  <mmd:activity_type>Numerical Simulation</mmd:activity_type>
  <mmd:dataset_citation>
    <mmd:author>Senter for utvikling av værvarslingstjenesten</mmd:author>
    <mmd:publication_date>2023-08-10T09:54:58Z</mmd:publication_date>
    <mmd:title>Arome-Arctic 2.5Km deterministic 2023-08-10T06:00:00Z + 66 hours</mmd:title>
  </mmd:dataset_citation>
</mmd:mmd>
`

var MMDExample = MetadataMMD{
	Title: &MMDTitle{
		Title:    "Arome Arctic",
		Language: "en",
	},
	DataCenter: &MMDDataCenter{
		DataCenterName: MMDShortLongName{
			ShortName: "MET Norway",
			LongName:  "Norwegian Meteorological Institute",
		},
		DataCenterURL: "https://data.met.no",
	},
}
