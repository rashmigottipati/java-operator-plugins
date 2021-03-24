package templates

import "sigs.k8s.io/kubebuilder/v3/pkg/model/file"

var _ file.Template = &PomXmlFile{}

type PomXmlFile struct {
	file.TemplateMixin
}

func (f *PomXmlFile) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "pom.xml"
	}

	f.TemplateBody = pomxmlTemplate

	return nil
}

// TODO: pass in the name of the operator i.e. replace Memcached
const pomxmlTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <parent>
    <artifactId>quarkus-operator-sdk-parent</artifactId>
    <groupId>io.quarkiverse.operatorsdk</groupId>
    <version>1.8.1-SNAPSHOT</version>
  </parent>
  <modelVersion>4.0.0</modelVersion>

  <artifactId>quarkus-operator-sdk-samples</artifactId>
  <name>Quarkus - Operator SDK - Samples</name>
  <packaging>jar</packaging>

  <properties>
    <maven.compiler.parameters>true</maven.compiler.parameters>
  </properties>

  <dependencies>
    <dependency>
      <groupId>io.quarkiverse.operatorsdk</groupId>
      <artifactId>quarkus-operator-sdk</artifactId>
      <version>${project.version}</version>
    </dependency>
    <dependency>
      <groupId>io.fabric8</groupId>
      <artifactId>crd-generator-apt</artifactId>
      <version>${kubernetes-client.version}</version>
    </dependency>
    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-resteasy-jackson</artifactId>
    </dependency>
    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-rest-client</artifactId>
    </dependency>
    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-rest-client-jackson</artifactId>
    </dependency>


  </dependencies>

  <build>
    <plugins>
      <plugin>
        <groupId>io.quarkus</groupId>
        <artifactId>quarkus-maven-plugin</artifactId>
        <executions>
          <execution>
            <goals>
              <goal>build</goal>
            </goals>
          </execution>
        </executions>
      </plugin>
    </plugins>
  </build>

</project>
`

