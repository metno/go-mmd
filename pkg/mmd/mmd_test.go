package mmd

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMMDEncoding(t *testing.T) {
	mmdDoc := &MetadataMMD{
		Title: MMDTitle{
			Title:    "Test",
			Language: "en",
		},
		Abstract: MMDAbstract{
			Abstract: "Testing, testing, testing.",
			Language: "en",
		},
		LastMetadataUpdate: "2020-01-20T08:00:00Z",
		Keywords: MMDKeywords{
			Keyword:    []string{"hepp", "blurgh"},
			Vocabulary: "mmd",
		},
		Collections:        []string{"NMDC", "SIOS"},
		MetadataIdentifier: "cfe3eb1a-4362-11ea-9f32-8fd0be8c935b",
		MetadataStatus:     "Active",
		GeographicExtent: GeographicExtent{
			Rectangle: Rectangle{
				North: 180,
				South: 90,
				West:  20,
				East:  10,
			},
		},
		DataCenter: MMDDataCenter{
			DataCenterName: MMDShortLongName{
				ShortName: "MET",
				LongName:  "Super long name MET",
			},
		},
		DataAccess: MMDDataAccess{
			Type:      "OGC",
			WMSLayers: []string{"Heppers", "blurgh"},
		},
		ISOTopicCategories: []string{"ocean"},
	}
	payload, err := xml.Marshal(mmdDoc)
	if err != nil {
		t.Errorf("Failed to seralize mmd: %s", err)
	}

	fmt.Println(string(payload))
}

// func TestMMDDecoding(t *testing.T) {
// 	var metadata MetadataMMD
// 	err := xml.Unmarshal([]byte(mmdExample), &metadata)
// 	if err != nil {
// 		t.Errorf("Expected successfull decoding of MMD:\n got error: %s\n wanted: no error", err)
// 	}
// }

var mmdExample = `
<?xml version="1.0"?>
<mmd:mmd xmlns:mmd="http://www.met.no/schema/mmd">
  <mmd:metadata_identifier>e8b2e2f5-b469-5ffa-986f-cd3e336db5b1</mmd:metadata_identifier>
  <mmd:metadata_status>active</mmd:metadata_status>
  <mmd:collection>ADC</mmd:collection>
  <mmd:collection>SIOS</mmd:collection>
  <mmd:collection>CC</mmd:collection>
  <mmd:collection>NSDN</mmd:collection>
  <mmd:title xml:lang="en">Sea Ice Extent</mmd:title>
  <mmd:abstract xml:lang="en">Arctic Sea Ice Extent estimates based upon EUMETSAT Ocean and Sea Ice SAF 
products. Both reprocessed and operational datasets have been used. See 
http://cryoclim.met.no/indicators/seaiceextent/ for details. </mmd:abstract>
  <mmd:related_information>
    <mmd:type>Dataset landing page</mmd:type>
    <mmd:description/>
    <mmd:resource>http://thredds.met.no/thredds/catalog/metusers/steingod/catalog.html?dataset=metusers/steingod/ci-sie-nh_osisaf_monthly_mean_sie.nc</mmd:resource>
  </mmd:related_information>
  <mmd:last_metadata_update>2013-10-30T08:09:59Z</mmd:last_metadata_update>
  <mmd:dataset_production_status>Complete</mmd:dataset_production_status>
  <mmd:dataset_language>en</mmd:dataset_language>
  <mmd:iso_topic_category>climatologyMeteorologyAtmosphere</mmd:iso_topic_category>
  <mmd:iso_topic_category>oceans</mmd:iso_topic_category>
  <mmd:project>
    <mmd:short_name/>
    <mmd:long_name>CryoClim</mmd:long_name>
  </mmd:project>
  <mmd:activity_type>Climate indicator</mmd:activity_type>
  <mmd:platform>
    <mmd:short_name/>
    <mmd:long_name>Not available</mmd:long_name>
  </mmd:platform>
  <?xml version="1.0"?>
  <mmd:mmd xmlns:mmd="http://www.met.no/schema/mmd">
	<mmd:metadata_identifier>e8b2e2f5-b469-5ffa-986f-cd3e336db5b1</mmd:metadata_identifier>
	<mmd:metadata_status>active</mmd:metadata_status>
	<mmd:collection>ADC</mmd:collection>
	<mmd:collection>SIOS</mmd:collection>
	<mmd:collection>CC</mmd:collection>
	<mmd:collection>NSDN</mmd:collection>
	<mmd:title xml:lang="en">Sea Ice Extent</mmd:title>
	<mmd:abstract xml:lang="en">Arctic Sea Ice Extent estimates based upon EUMETSAT Ocean and Sea Ice SAF 
  products. Both reprocessed and operational datasets have been used. See 
  http://cryoclim.met.no/indicators/seaiceextent/ for details. </mmd:abstract>
	<mmd:related_information>
	  <mmd:type>Dataset landing page</mmd:type>
	  <mmd:description/>
	  <mmd:resource>http://thredds.met.no/thredds/catalog/metusers/steingod/catalog.html?dataset=metusers/steingod/ci-sie-nh_osisaf_monthly_mean_sie.nc</mmd:resource>
	</mmd:related_information>
	<mmd:last_metadata_update>2013-10-30T08:09:59Z</mmd:last_metadata_update>
	<mmd:dataset_production_status>Complete</mmd:dataset_production_status>
	<mmd:dataset_language>en</mmd:dataset_language>
	<mmd:iso_topic_category>climatologyMeteorologyAtmosphere</mmd:iso_topic_category>
	<mmd:iso_topic_category>oceans</mmd:iso_topic_category>
	<mmd:project>
	  <mmd:short_name/>
	  <mmd:long_name>CryoClim</mmd:long_name>
	</mmd:project>
	<mmd:activity_type>Climate indicator</mmd:activity_type>
	<mmd:platform>
	  <mmd:short_name/>
	  <mmd:long_name>Not available</mmd:long_name>
	</mmd:platform>
	<mmd:keywords vocabulary="CF">
    <mmd:keyword>Climate</mmd:keyword>
    <mmd:keyword>Sea Ice</mmd:keyword>
    <mmd:keyword>Oceanography</mmd:keyword>
    <mmd:keyword>Sea Ice Extent</mmd:keyword>
    <mmd:keyword>Meteorology</mmd:keyword>
    <mmd:keyword>Remote Sensing</mmd:keyword>
  </mmd:keywords>
  <mmd:keywords vocabulary="gcmd">
    <mmd:keyword>Cryosphere &gt; Sea Ice &gt; Sea Ice Extent</mmd:keyword>
    <mmd:keyword>Geographic Region &gt; Northern Hemisphere</mmd:keyword>
    <mmd:keyword>Oceans &gt; Sea Ice &gt; Sea Ice Extent</mmd:keyword>
    <mmd:keyword>Vertical Location &gt; Sea Surface</mmd:keyword>
    <mmd:keyword>EUMETSAT/OSISAF &gt; Satellite Application Facility on Ocean and Sea Ice, European Organisation for the Exploitation of Meteorological Satellites</mmd:keyword>
  </mmd:keywords>
</mmd:mmd>
`
