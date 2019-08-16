package model_test

import (
	"encoding/xml"
	"testing"

	"github.com/devblok/koru/model"
)

var Cube_file string = `
<?xml version="1.0" encoding="utf-8"?>
<COLLADA xmlns="http://www.collada.org/2005/11/COLLADASchema" version="1.4.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <asset>
    <contributor>
      <author>Blender User</author>
      <authoring_tool>Blender 2.79.0 commit date:1970-01-01, commit time:00:00, hash:unknown</authoring_tool>
    </contributor>
    <created>2019-07-31T20:40:46</created>
    <modified>2019-07-31T20:40:46</modified>
    <unit name="meter" meter="1"/>
    <up_axis>Z_UP</up_axis>
  </asset>
  <library_images/>
  <library_effects>
    <effect id="Material-effect">
      <profile_COMMON>
        <technique sid="common">
          <phong>
            <emission>
              <color sid="emission">0 0 0 1</color>
            </emission>
            <ambient>
              <color sid="ambient">0 0 0 1</color>
            </ambient>
            <diffuse>
              <color sid="diffuse">0.64 0.64 0.64 1</color>
            </diffuse>
            <specular>
              <color sid="specular">0.5 0.5 0.5 1</color>
            </specular>
            <shininess>
              <float sid="shininess">50</float>
            </shininess>
            <index_of_refraction>
              <float sid="index_of_refraction">1</float>
            </index_of_refraction>
          </phong>
        </technique>
      </profile_COMMON>
    </effect>
  </library_effects>
  <library_materials>
    <material id="Material-material" name="Material">
      <instance_effect url="#Material-effect"/>
    </material>
  </library_materials>
  <library_geometries>
    <geometry id="Cube-mesh" name="Cube">
      <mesh>
        <source id="Cube-mesh-positions">
          <float_array id="Cube-mesh-positions-array" count="24">1 1 -1 1 -1 -1 -1 -0.9999998 -1 -0.9999997 1 -1 1 0.9999995 1 0.9999994 -1.000001 1 -1 -0.9999997 1 -1 1 1</float_array>
          <technique_common>
            <accessor source="#Cube-mesh-positions-array" count="8" stride="3">
              <param name="X" type="float"/>
              <param name="Y" type="float"/>
              <param name="Z" type="float"/>
            </accessor>
          </technique_common>
        </source>
        <source id="Cube-mesh-normals">
          <float_array id="Cube-mesh-normals-array" count="36">0 0 -1 0 0 1 1 0 -2.38419e-7 0 -1 -4.76837e-7 -1 2.38419e-7 -1.49012e-7 2.68221e-7 1 2.38419e-7 0 0 -1 0 0 1 1 -5.96046e-7 3.27825e-7 -4.76837e-7 -1 0 -1 2.38419e-7 -1.19209e-7 2.08616e-7 1 0</float_array>
          <technique_common>
            <accessor source="#Cube-mesh-normals-array" count="12" stride="3">
              <param name="X" type="float"/>
              <param name="Y" type="float"/>
              <param name="Z" type="float"/>
            </accessor>
          </technique_common>
        </source>
        <vertices id="Cube-mesh-vertices">
          <input semantic="POSITION" source="#Cube-mesh-positions"/>
        </vertices>
        <triangles material="Material-material" count="12">
          <input semantic="VERTEX" source="#Cube-mesh-vertices" offset="0"/>
          <input semantic="NORMAL" source="#Cube-mesh-normals" offset="1"/>
          <p>0 0 2 0 3 0 7 1 5 1 4 1 4 2 1 2 0 2 5 3 2 3 1 3 2 4 7 4 3 4 0 5 7 5 4 5 0 6 1 6 2 6 7 7 6 7 5 7 4 8 5 8 1 8 5 9 6 9 2 9 2 10 6 10 7 10 0 11 3 11 7 11</p>
        </triangles>
      </mesh>
    </geometry>
  </library_geometries>
  <library_controllers/>
  <library_visual_scenes>
    <visual_scene id="Scene" name="Scene">
      <node id="Cube" name="Cube" type="NODE">
        <matrix sid="transform">1 0 0 0 0 1 0 0 0 0 1 0 0 0 0 1</matrix>
        <instance_geometry url="#data, ioutil.ReadAllCube-mesh" name="Cube">
          <bind_material>
            <technique_common>
              <instance_material symbol="Material-material" target="#Material-material"/>
            </technique_common>
          </bind_material>
        </instance_geometry>
      </node>
    </visual_scene>
  </library_visual_scenes>
  <scene>
    <instance_visual_scene url="#Scene"/>
  </scene>
</COLLADA>
`

var Cube_file_wTex = `
<?xml version="1.0" encoding="utf-8"?>
<COLLADA xmlns="http://www.collada.org/2005/11/COLLADASchema" version="1.4.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <asset>
    <contributor>
      <author>Blender User</author>
      <authoring_tool>Blender 2.80.75 commit date:2019-07-29, commit time:14:47, hash:f6cb5f54494e</authoring_tool>
    </contributor>
    <created>2019-08-06T13:32:51</created>
    <modified>2019-08-06T13:32:51</modified>
    <unit name="meter" meter="1"/>
    <up_axis>Z_UP</up_axis>
  </asset>
  <library_cameras>
    <camera id="Camera-camera" name="Camera">
      <optics>
        <technique_common>
          <perspective>
            <xfov sid="xfov">39.59775</xfov>
            <aspect_ratio>1.777778</aspect_ratio>
            <znear sid="znear">0.1</znear>
            <zfar sid="zfar">100</zfar>
          </perspective>
        </technique_common>
      </optics>
      <extra>
        <technique profile="blender">
          <shiftx sid="shiftx" type="float">0</shiftx>
          <shifty sid="shifty" type="float">0</shifty>
          <dof_distance sid="dof_distance" type="float">10</dof_distance>
        </technique>
      </extra>
    </camera>
  </library_cameras>
  <library_lights>
    <light id="Light-light" name="Light">
      <technique_common>
        <point>
          <color sid="color">1000 1000 1000</color>
          <constant_attenuation>1</constant_attenuation>
          <linear_attenuation>0</linear_attenuation>
          <quadratic_attenuation>0.00111109</quadratic_attenuation>
        </point>
      </technique_common>
      <extra>
        <technique profile="blender">
          <type sid="type" type="int">0</type>
          <flag sid="flag" type="int">0</flag>
          <mode sid="mode" type="int">1</mode>
          <gamma sid="blender_gamma" type="float">1</gamma>
          <red sid="red" type="float">1</red>
          <green sid="green" type="float">1</green>
          <blue sid="blue" type="float">1</blue>
          <shadow_r sid="blender_shadow_r" type="float">0</shadow_r>
          <shadow_g sid="blender_shadow_g" type="float">0</shadow_g>
          <shadow_b sid="blender_shadow_b" type="float">0</shadow_b>
          <energy sid="blender_energy" type="float">1000</energy>
          <dist sid="blender_dist" type="float">29.99998</dist>
          <spotsize sid="spotsize" type="float">75</spotsize>
          <spotblend sid="spotblend" type="float">0.15</spotblend>
          <att1 sid="att1" type="float">0</att1>
          <att2 sid="att2" type="float">1</att2>
          <falloff_type sid="falloff_type" type="int">2</falloff_type>
          <clipsta sid="clipsta" type="float">0.04999995</clipsta>
          <clipend sid="clipend" type="float">30.002</clipend>
          <bias sid="bias" type="float">1</bias>
          <soft sid="soft" type="float">3</soft>
          <bufsize sid="bufsize" type="int">2880</bufsize>
          <samp sid="samp" type="int">3</samp>
          <buffers sid="buffers" type="int">1</buffers>
          <area_shape sid="area_shape" type="int">1</area_shape>
          <area_size sid="area_size" type="float">0.1</area_size>
          <area_sizey sid="area_sizey" type="float">0.1</area_sizey>
          <area_sizez sid="area_sizez" type="float">1</area_sizez>
        </technique>
      </extra>
    </light>
  </library_lights>
  <library_effects>
    <effect id="Material-effect">
      <profile_COMMON>
        <technique sid="common">
          <lambert>
            <emission>
              <color sid="emission">0 0 0 1</color>
            </emission>
            <diffuse>
              <color sid="diffuse">0.8 0.8 0.8 1</color>
            </diffuse>
            <index_of_refraction>
              <float sid="ior">1.45</float>
            </index_of_refraction>
          </lambert>
        </technique>
      </profile_COMMON>
    </effect>
  </library_effects>
  <library_images/>
  <library_materials>
    <material id="Material-material" name="Material">
      <instance_effect url="#Material-effect"/>
    </material>
  </library_materials>
  <library_geometries>
    <geometry id="Cube-mesh" name="Cube">
      <mesh>
        <source id="Cube-mesh-positions">
          <float_array id="Cube-mesh-positions-array" count="24">1 1 1 1 1 -1 1 -1 1 1 -1 -1 -1 1 1 -1 1 -1 -1 -1 1 -1 -1 -1</float_array>
          <technique_common>
            <accessor source="#Cube-mesh-positions-array" count="8" stride="3">
              <param name="X" type="float"/>
              <param name="Y" type="float"/>
              <param name="Z" type="float"/>
            </accessor>
          </technique_common>
        </source>
        <source id="Cube-mesh-normals">
          <float_array id="Cube-mesh-normals-array" count="18">0 0 1 0 -1 0 -1 0 0 0 0 -1 1 0 0 0 1 0</float_array>
          <technique_common>
            <accessor source="#Cube-mesh-normals-array" count="6" stride="3">
              <param name="X" type="float"/>
              <param name="Y" type="float"/>
              <param name="Z" type="float"/>
            </accessor>
          </technique_common>
        </source>
        <source id="Cube-mesh-map-0">
          <float_array id="Cube-mesh-map-0-array" count="72">0.625 0 0.375 0.25 0.375 0 0.625 0.25 0.375 0.5 0.375 0.25 0.625 0.5 0.375 0.75 0.375 0.5 0.625 0.75 0.375 1 0.375 0.75 0.375 0.5 0.125 0.75 0.125 0.5 0.875 0.5 0.625 0.75 0.625 0.5 0.625 0 0.625 0.25 0.375 0.25 0.625 0.25 0.625 0.5 0.375 0.5 0.625 0.5 0.625 0.75 0.375 0.75 0.625 0.75 0.625 1 0.375 1 0.375 0.5 0.375 0.75 0.125 0.75 0.875 0.5 0.875 0.75 0.625 0.75</float_array>
          <technique_common>
            <accessor source="#Cube-mesh-map-0-array" count="36" stride="2">
              <param name="S" type="float"/>
              <param name="T" type="float"/>
            </accessor>
          </technique_common>
        </source>
        <vertices id="Cube-mesh-vertices">
          <input semantic="POSITION" source="#Cube-mesh-positions"/>
        </vertices>
        <triangles material="Material-material" count="12">
          <input semantic="VERTEX" source="#Cube-mesh-vertices" offset="0"/>
          <input semantic="NORMAL" source="#Cube-mesh-normals" offset="1"/>
          <input semantic="TEXCOORD" source="#Cube-mesh-map-0" offset="2" set="0"/>
          <p>4 0 0 2 0 1 0 0 2 2 1 3 7 1 4 3 1 5 6 2 6 5 2 7 7 2 8 1 3 9 7 3 10 5 3 11 0 4 12 3 4 13 1 4 14 4 5 15 1 5 16 5 5 17 4 0 18 6 0 19 2 0 20 2 1 21 6 1 22 7 1 23 6 2 24 4 2 25 5 2 26 1 3 27 3 3 28 7 3 29 0 4 30 2 4 31 3 4 32 4 5 33 0 5 34 1 5 35</p>
        </triangles>
      </mesh>
    </geometry>
  </library_geometries>
  <library_visual_scenes>
    <visual_scene id="Scene" name="Scene">
      <node id="Camera" name="Camera" type="NODE">
        <matrix sid="transform">0.6859207 -0.3240135 0.6515582 7.358891 0.7276763 0.3054208 -0.6141704 -6.925791 0 0.8953956 0.4452714 4.958309 0 0 0 1</matrix>
        <instance_camera url="#Camera-camera"/>
      </node>
      <node id="Light" name="Light" type="NODE">
        <matrix sid="transform">-0.2908646 -0.7711008 0.5663932 4.076245 0.9551712 -0.1998834 0.2183912 1.005454 -0.05518906 0.6045247 0.7946723 5.903862 0 0 0 1</matrix>
        <instance_light url="#Light-light"/>
      </node>
      <node id="Cube" name="Cube" type="NODE">
        <matrix sid="transform">1 0 0 0 0 1 0 0 0 0 1 0 0 0 0 1</matrix>
        <instance_geometry url="#Cube-mesh" name="Cube">
          <bind_material>
            <technique_common>
              <instance_material symbol="Material-material" target="#Material-material">
                <bind_vertex_input semantic="UVMap" input_semantic="TEXCOORD" input_set="0"/>
              </instance_material>
            </technique_common>
          </bind_material>
        </instance_geometry>
      </node>
    </visual_scene>
  </library_visual_scenes>
  <scene>
    <instance_visual_scene url="#Scene"/>
  </scene>
</COLLADA>
`

func TestImportColladaObject(t *testing.T) {
	obj, err := model.ImportColladaObject([]byte(Cube_file), nil)
	if err != nil {
		t.Fatal(err)
	}

	vert := obj.Vertices()
	if len(vert) != 36 {
		t.Fatalf("wrong amount of vertices, got: %d", len(vert))
	}
}

func TestImportColladaObjectWTex(t *testing.T) {
	obj, err := model.ImportColladaObject([]byte(Cube_file_wTex), nil)
	if err != nil {
		t.Fatal(err)
	}

	vert := obj.Vertices()
	if len(vert) != 36 {
		t.Fatalf("wrong amount of vertices, got: %d", len(vert))
	}
}

func TestTrianglesDecode(t *testing.T) {
	data := `
		<triangles material="Material-material" count="12">
		<input semantic="VERTEX" source="#Cube-mesh-vertices" offset="0"/>
		<input semantic="NORMAL" source="#Cube-mesh-normals" offset="1"/>
		<p>0 0 2 0 3 0 7 1 5 1 4 1 4 2 1 2 0 2 5 3 2 3 1 3 2 4 7 4 3 4 0 5 7 5 4 5 0 6 1 6 2 6 7 7 6 7 5 7 4 8 5 8 1 8 5 9 6 9 2 9 2 10 6 10 7 10 0 11 3 11 7 11</p>
		</triangles>
	`
	var triangles model.Triangles
	err := xml.Unmarshal([]byte(data), &triangles)
	if err != nil {
		t.Fatal(err)
	}

	if triangles.Material != "Material-material" {
		t.Fatalf("incorrect material: %s", triangles.Material)
	}

	if triangles.Count != 12 {
		t.Fatalf("incorrect count: %d", triangles.Count)
	}

	if len(triangles.Inputs) != 2 {
		t.Fatalf("number of inputs incorrect: %d", len(triangles.Inputs))
	}

	if len(triangles.Index) != 12*6 {
		t.Fatalf("number of index elements incorrect: %d", len(triangles.Index))
	}
}

func TestInputDecode(t *testing.T) {
	data := `
	<object>
		<input semantic="VERTEX" source="#Cube-mesh-vertices" offset="0" />
		<input semantic="NORMAL" source="#Cube-mesh-normals" offset="1" />
		<input semantic="TEXTUR" source="#Cube-mesh-textures" offset="2" />
	</object>
	`

	type Object struct {
		XMLNname xml.Name      `xml:"object"`
		Inputs   []model.Input `xml:"input"`
	}

	var obj Object
	err := xml.Unmarshal([]byte(data), &obj)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if obj.Inputs[0].Offset != 0 || obj.Inputs[0].Semantic != "VERTEX" || obj.Inputs[0].Source != "#Cube-mesh-vertices" {
		t.Logf("expected Offset: %d, got: %d\nexpected Semantic: %s, got %s\nexpected Source: %s, got: %s", 0, obj.Inputs[0].Offset, "VERTEX", obj.Inputs[0].Semantic, "#Cube-mesh-vertices", obj.Inputs[0].Source)
		t.Fail()
	}
	if obj.Inputs[1].Offset != 1 || obj.Inputs[1].Semantic != "NORMAL" || obj.Inputs[1].Source != "#Cube-mesh-normals" {
		t.Logf("expected Offset: %d, got: %d\nexpected Semantic: %s, got %s\nexpected Source: %s, got: %s", 1, obj.Inputs[1].Offset, "NORMAL", obj.Inputs[1].Semantic, "#Cube-mesh-normals", obj.Inputs[1].Source)
		t.Fail()
	}
	if obj.Inputs[2].Offset != 2 || obj.Inputs[2].Semantic != "TEXTUR" || obj.Inputs[2].Source != "#Cube-mesh-textures" {
		t.Logf("expected Offset: %d, got: %d\nexpected Semantic: %s, got %s\nexpected Source: %s, got: %s", 2, obj.Inputs[2].Offset, "TEXTUR", obj.Inputs[2].Semantic, "#Cube-mesh-textures", obj.Inputs[2].Source)
		t.Fail()
	}
}

func TestFloatsDecode(t *testing.T) {
	data := `<float_array id="Cube-mesh-normals-array" count="36">0 0 -1 0 0 1 1 0 -2.38419e-7 0 -1 -4.76837e-7 -1 2.38419e-7 -1.49012e-7 2.68221e-7 1 2.38419e-7 0 0 -1 0 0 1 1 -5.96046e-7 3.27825e-7 -4.76837e-7 -1 0 -1 2.38419e-7 -1.19209e-7 2.08616e-7 1 0</float_array>`

	var floats model.Floats
	if err := xml.Unmarshal([]byte(data), &floats); err != nil {
		t.Fatal(err)
	}

	if len(floats.Data) != 36 {
		t.Fatalf("bad number of floats, got: %d", len(floats.Data))
	}

	if floats.ID != "Cube-mesh-normals-array" {
		t.Fatalf("bad id, got: %s", floats.ID)
	}
}
