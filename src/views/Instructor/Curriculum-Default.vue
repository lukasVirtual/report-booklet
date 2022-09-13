<template>
  <base-layout v-if="$route.name === 'Curriculum'">
    <v-main style="display: flex; margin: auto">
      <v-select
        :items="who"
        variant="underlined"
        style="max-width: 200px; max-height: 60px"
        v-model="selected"
      >
      </v-select>
      <v-table
        style="width: 1100px; max-width: 1400px; border-radius: 10px"
        class="mt-5"
      >
        <tbody>
          <tr v-for="elem in items" :key="elem.title">
            <td>{{ elem.title }}</td>
            <td class="d-flex">
              <v-checkbox
                color="primary"
                class="align-center justify-center"
              ></v-checkbox>
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler";
import { defineComponent, ref, onMounted } from "@vue/runtime-core";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "CurriculumDefault",
  components: { BaseLayout },
  setup() {
    const selected = ref("");
    const who = ref<string[]>([]);

    onMounted(async () => {
      const instructor = await loginService.getUser();
      const users = await dataService.GetAllUserForInstructor(instructor);
      who.value = users.map((v: { Name: string }) => v.Name);
    });

    const items = [
      {
        title: "Grundsätze und Methoden des Projektmanagements anwenden",
        value: 1,
      },
      {
        title:
          "Auftragsunterlagen und Durchführbarkeit des Auftrags prüfen, insbesondere in Hinblick auf rechtliche, wirtschaftliche und terminliche Vorgaben, und den Auftrag mit den betrieblichen Prozessen und Möglichkeiten abstimmen",
        value: 2,
      },
      {
        title:
          "Zeitplan und Reihenfolge der Arbeitsschritte für den eigenen Arbeitsbereich festlegen",
        value: 3,
      },
      {
        title:
          "Termine planen und abstimmen sowie Terminüberwachung durchführen",
        value: 4,
      },
      {
        title:
          "Probleme analysieren und als Aufgabe definieren sowie Lösungsalternativen entwickeln und beurteilen",
        value: 4,
      },
      {
        title:
          "Arbeits- und Organisationsmittel wirtschaftlich und ökologisch unter Berücksichtigung der vorhandenen Ressourcen und der Budgetvorgaben einsetzen",
        value: 4,
      },
      {
        title:
          "Aufgaben im Team sowie mit internen und externen Kunden und Kundinnen planen und abstimmen",
        value: 4,
      },
      {
        title:
          "betriebswirtschaftlich relevante Daten erheben und bewerten und dabei Geschäfts- und Leistungsprozesse berücksichtigen",
        value: 4,
      },
      {
        title:
          "eigene Vorgehensweise sowie die Aufgabendurchführung im Team reflektieren und bei der Verbesserung der Arbeitsprozesse mitwirken",
        value: 4,
      },
      {
        title:
          "im Rahmen der Marktbeobachtung Preise, Leistungen und Konditionen von Wettbewerbern vergleichen",
        value: 4,
      },
      {
        title:
          "Bedarfe von Kunden und Kundinnen feststellen sowie Zielgruppen unterscheiden",
        value: 4,
      },
      {
        title:
          "Kunden unter Beachtung von Kommunikationsregeln informieren sowie Sachverhalte präsentieren und dabei deutsche und englische Fachbegriffe anwenden",
        value: 4,
      },
      {
        title: "Maßnahmen für Marketing und Vertrieb unterstützen",
        value: 4,
      },
      {
        title:
          "Informationsquellen auch in englischer Sprache aufgabenbezogen auswerten und für die Kundeninformation nutzen",
        value: 4,
      },
      {
        title:
          "marktgängige IT-Systeme für unterschiedliche Einsatzbereiche hinsichtlich Leistungsfähigkeit, Wirtschaftlichkeit und Barrierefreiheit beurteilen",
        value: 4,
      },
      {
        title:
          "Angebote zu IT-Komponenten, IT-Produkten und IT-Dienstleistungen einholen und bewerten sowie Spezifikationen und Konditionen vergleichen",
        value: 4,
      },
      {
        title:
          "IT -Systeme zur Bearbeitung betrieblicher Fachaufgaben analysieren sowie unter Beachtung insbesondere von Lizenzmodellen und Urheberrechten und Barrierefreiheit konzeptionieren, konfigurieren, testen und dokumentieren",
        value: 4,
      },
      {
        title:
          "Programmiersprachen, insbesondere prozedurale und objektorientierte Programmiersprachen, unterscheiden",
        value: 4,
      },
      {
        title:
          "betriebliche Qualitätssicherungssysteme im eigenen Arbeitsbereich anwenden und Qualitätssicherungsmaßnahmen projektbegleitend durchführen und dokumentieren",
        value: 4,
      },
      {
        title:
          "betriebliche Vorgaben und rechtliche Regelungen zur IT-Sicherheit und zum Datenschutz einhalten",
        value: 4,
      },
      {
        title:
          "Sicherheitsanforderungen von IT -Systemen analysieren und Maßnahmen zur IT - Sicherheit ableiten, abstimmen, umsetzen und evaluieren",
        value: 4,
      },
      {
        title:
          "Leistungen nach betrieblichen und vertraglichen Vorgaben dokumentieren",
        value: 4,
      },
      {
        title:
          "Leistungserbringung unter Berücksichtigung der organisatorischen und terminlichen Vorgaben mit Kunden und Kundinnen abstimmen und kontrollieren",
        value: 4,
      },
      {
        title: "Veränderungsprozesse begleiten und unterstützen",
        value: 4,
      },
      {
        title:
          "Kunden und Kundinnen in die Nutzung von Produkten und Dienstleistungen einweisen",
        value: 4,
      },
      {
        title:
          "Leistungen und Dokumentationen an Kunden und Kundinnen übergeben sowie Abnahmeprotokolle anfertigen",
        value: 4,
      },
      {
        title:
          "Kosten für erbrachte Leistungen erfassen sowie im Zeitvergleich und im Soll-IstVergleich bewerten",
        value: 4,
      },
      {
        title:
          "Netzwerkkonzepte für unterschiedliche Anwendungsgebiete unterscheiden",
        value: 4,
      },
      {
        title: "Datenaustausch von vernetzten Systemen realisieren",
        value: 4,
      },
      {
        title:
          "Verfügbarkeit und Ausfallwahrscheinlichkeiten analysieren und Lösungsvorschläge unterbreiten",
        value: 4,
      },
      {
        title:
          "Maßnahmen zur präventiven Wartung und zur Störungsvermeidung einleiten und durchführen",
        value: 4,
      },
      {
        title:
          "Programmspezifikationen festlegen, Datenmodelle und Strukturen aus fachlichen Anforderungen ableiten sowie Schnittstellen festlegen",
        value: 4,
      },
      {
        title:
          "Programmiersprachen auswählen und unterschiedliche Programmiersprachen anwenden",
        value: 4,
      },
      {
        title:
          "Vorgehensmodelle und -methoden sowie Entwicklungsumgebungen und bibliotheken auswählen und einsetzen",
        value: 4,
      },
      {
        title: "Analyse- und Designverfahren anwenden",
        value: 4,
      },
      {
        title:
          "Benutzerschnittstellen ergonomisch gestalten und an Kundenanforderungen anpassen",
        value: 4,
      },
      {
        title:
          "Sicherheitsaspekte bei der Entwicklung von Softwareanwendungen berücksichtigen",
        value: 4,
      },
      {
        title: "Datenintegrität mithilfe von Werkzeugen sicherstellen",
        value: 4,
      },
      {
        title: "Modultests erstellen und durchführen",
        value: 4,
      },
      {
        title:
          "gegenseitige Wertschätzung unter Berücksichtigung gesellschaftlicher Vielfalt bei betrieblichen Abläufen praktizieren",
        value: 4,
      },
      {
        title:
          "Strategien zum verantwortungsvollen Umgang mit digitalen Medien anwenden und im virtuellen Raum unter Wahrung der Persönlichkeitsrechte Dritter zusammenarbeiten",
        value: 4,
      },
      {
        title:
          "insbesondere bei der Speicherung, Darstellung und Weitergabe digitaler Inhalte die Auswirkungen des eigenen Kommunikations- und Informationsverhaltens berücksichtigen",
        value: 4,
      },
      {
        title:
          "bei der Beurteilung, Entwicklung, Umsetzung und Betreuung von IT-Lösungen ethische Aspekte reflektieren",
        value: 4,
      },
      {
        title:
          "Gespräche situationsgerecht führen und Kunden und Kundinnen unter Berücksichtigung der Kundeninteressen beraten",
        value: 1,
      },
      {
        title:
          "Kundenbeziehungen unter Beachtung rechtlicher Regelungen und betrieblicher Grundsätze gestalten",
        value: 1,
      },
      {
        title:
          "Daten und Sachverhalte interpretieren, multimedial aufbereiten und situationsgerecht unter Nutzung digitaler Werkzeuge und unter Berücksichtigung der betrieblichen Vorgaben präsentieren",
        value: 1,
      },
      {
        title:
          "technologische Entwicklungstrends von IT-Systemen feststellen sowie ihre wirtschaftlichen, sozialen und beruflichen Auswirkungen aufzeigen",
        value: 1,
      },
      {
        title:
          "Veränderungen von Einsatzfeldern für ITSysteme aufgrund technischer, wirtschaftlicher und gesellschaftlicher Entwicklungen feststellen",
        value: 1,
      },
      {
        title: "systematisch Fehler erkennen, analysieren und beheben",
        value: 1,
      },
      {
        title:
          "Algorithmen formulieren und Anwendungen in einer Programmiersprache erstellen",
        value: 1,
      },
      {
        title:
          "Datenbankmodelle unterscheiden, Daten organisieren und speichern sowie Abfragen erstellen",
        value: 1,
      },
      {
        title:
          "Ursachen von Qualitätsmängeln systematisch feststellen, beseitigen und dokumentieren",
        value: 1,
      },
      {
        title:
          "im Rahmen eines Verbesserungsprozesses die Zielerreichung kontrollieren, insbesondere einen Soll-Ist-Vergleich durchführen",
        value: 1,
      },
      {
        title:
          "Bedrohungsszenarien erkennen und Schadenspotenziale unter Berücksichtigung wirtschaftlicher und technischer Kriterien einschätzen",
        value: 1,
      },
      {
        title:
          "Kunden und Kundinnen im Hinblick auf Anforderungen an die IT-Sicherheit und an den Datenschutz beraten",
        value: 1,
      },
      {
        title:
          "Wirksamkeit und Effizienz der umgesetzten Maßnahmen zur IT-Sicherheit und zum Datenschutz prüfen",
        value: 1,
      },
      {
        title:
          "Störungsmeldungen aufnehmen und analysieren sowie Maßnahmen zur Störungsbeseitigung ergreifen",
        value: 1,
      },
      {
        title:
          "Dokumentationen zielgruppengerecht und barrierefrei anfertigen, bereitstellen und pflegen, insbesondere technische Dokumentationen, System - sowie Benutzerdokumentationen",
        value: 1,
      },
      {
        title:
          "Sicherheitsmechanismen, insbesondere Zugriffsmöglichkeiten und -rechte, festlegen und implementieren",
        value: 1,
      },
      {
        title: "Speicherlösungen, insbesondere Datenbanksysteme, integrieren",
        value: 1,
      },
      { title: "Teilaufgaben von IT-Systemen automatisieren", value: 1 },
      {
        title:
          "Anwendungslösungen unter Berücksichtigung der bestehenden Systemarchitektur entwerfen und realisieren",
        value: 1,
      },
      { title: "bestehende Anwendungslösungen anpassen", value: 1 },
      {
        title:
          "Datenaustausch zwischen Systemen realisieren und unterschiedliche Datenquellen nutzen",
        value: 1,
      },
      {
        title:
          "komplexe Abfragen aus unterschiedlichen Datenquellen durchführen und Datenbestandsberichte erstellen",
        value: 1,
      },
      { title: "Werkzeuge zur Versionsverwaltung einsetzen", value: 1 },
      {
        title:
          "Testkonzepte erstellen und Tests durchführen sowie Testergebnisse bewerten und dokumentieren",
        value: 1,
      },
      {
        title:
          "Daten und Sachverhalte aus Tests multimedial aufbereiten und situationsgerecht unter Nutzung digitaler Werkzeuge und unter Beachtung der betrieblichen Vorgaben präsentieren",
        value: 1,
      },
      {
        title:
          "wesentliche Inhalte und Bestandteile des Ausbildungsvertrages darstellen, Rechte und Pflichten aus dem Ausbildungsvertrag feststellen und Aufgaben der Beteiligten im dualen System beschreiben",
        value: 1,
      },
      {
        title:
          "den betrieblichen Ausbildungsplan mit der Ausbildungsordnung vergleichen",
        value: 1,
      },
      {
        title:
          "arbeits-, sozial- und mitbestimmungsrechtliche Vorschriften sowie für den Arbeitsbereich geltende Tarif- und Arbeitszeitregelungen beachten",
        value: 1,
      },
      { title: "Positionen der eigenen Entgeltabrechnung erklären", value: 1 },
      {
        title:
          "Chancen und Anforderungen des lebensbegleitenden Lernens für die berufliche und persönliche Entwicklung begründen und die eigenen Kompetenzen weiterentwickeln",
        value: 1,
      },
      {
        title:
          "Lern- und Arbeitstechniken sowie Methoden des selbstgesteuerten Lernens anwenden und beruflich relevante Informationsquellen nutzen",
        value: 1,
      },
      {
        title:
          "berufliche Aufstiegs- und Weiterentwicklungsmöglichkeiten darstellen",
        value: 1,
      },
      {
        title:
          "die Rechtsform und den organisatorischen Aufbau des Ausbildungsbetriebes mit seinen Aufgaben und Zuständigkeiten sowie die Zusammenhänge zwischen den Geschäftsprozessen erläutern",
        value: 1,
      },
      {
        title:
          "Beziehungen des Ausbildungsbetriebes und seiner Beschäftigten zu Wirtschaftsorganisationen, Berufsvertretungen und Gewerkschaften nennen",
        value: 1,
      },
      {
        title:
          "Grundlagen, Aufgaben und Arbeitsweise der betriebsverfassungsrechtlichen Organe des Ausbildungsbetriebes beschreiben",
        value: 1,
      },
      {
        title:
          "Gefährdung von Sicherheit und Gesundheit am Arbeitsplatz feststellen und Maßnahmen zur Vermeidung der Gefährdung ergreifen",
        value: 1,
      },
      {
        title:
          "berufsbezogene Arbeitsschutz- und Unfallverhütungsvorschriften anwenden",
        value: 1,
      },
      {
        title:
          "Verhaltensweisen bei Unfällen beschreiben sowie erste Maßnahmen einleiten",
        value: 1,
      },
      {
        title:
          "Vorschriften des vorbeugenden Brandschutzes anwenden sowie Verhaltensweisen bei Bränden beschreiben und Maßnahmen zur Brandbekämpfung ergreifen",
        value: 1,
      },
      {
        title:
          "mögliche Umweltbelastungen durch den Ausbildungsbetrieb und seinen Beitrag zum Umweltschutz an Beispielen erklären",
        value: 1,
      },
      {
        title:
          "für den Ausbildungsbetrieb geltende Regelungen des Umweltschutzes anwenden",
        value: 1,
      },
      {
        title:
          "Möglichkeiten der wirtschaftlichen und umweltschonenden Energie- und Materialverwendung nutzen",
        value: 1,
      },
      {
        title:
          "Abfälle vermeiden sowie Stoffe und Materialien einer umweltschonenden Entsorgung zuführen",
        value: 1,
      },
    ];

    return { items, selected, who };
  },
});
</script>
