<template>
  <div
    id="comp"
    v-if="$route.path === '/Berichtsheft'"
    @click="showDate(propsDate)"
  >
    <v-container fluid>
      <v-card
        style="
          width: 1100px;
          border: 2px solid black;
          overflow-x: hidden;
          border-radius: 10px;
        "
        tile
        class="mt-5 overflow-y-auto"
      >
        <v-card-actions style="margin: auto">
          <v-card-title>{{ propsDate }}</v-card-title>
          <v-spacer></v-spacer>
          <v-select
            :items="envItems"
            variant="underlined"
            style="max-width: 120px; max-height: 60px"
          ></v-select>
        </v-card-actions>
      </v-card>
      <v-card
        style="
          width: 1100px;
          height: 500px;
          border: 1px solid black;
          overflow-x: hidden;
          border-radius: 10px;
        "
        tile
        class="mt-4 overflow-y-auto"
      >
        <v-row dense>
          <v-col cols="1" class="text-center">
            <div
              style="
                background-color: #3e8ede;
                height: 320%;
                border-radius: 10px;
              "
            >
              <div style="height: 15px"></div>
              <v-btn @click="anyNumber++" icon elevation="20"
                ><v-icon size="30" color="blue">mdi-plus</v-icon></v-btn
              >
              <div style="height: 10px"></div>
              <v-btn icon elevation="20"
                ><v-icon size="30" color="red" @click="removeItems"
                  >mdi-trash-can</v-icon
                ></v-btn
              >
              <div style="height: 10px"></div>

              <qualifications-default></qualifications-default>
            </div>
          </v-col>
          <div style="width: 5px"></div>
          <v-row>
            <v-col cols="12" md="12">
              <div v-for="j in anyNumber" :key="j">
                <v-card-actions>
                  <typing-field></typing-field>
                </v-card-actions>
              </div>
            </v-col>
            <v-spacer></v-spacer>
            <v-col cols="2"> </v-col>
            <v-spacer></v-spacer>
          </v-row>
        </v-row>
      </v-card>
    </v-container>
  </div>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { defineComponent, ref } from "vue";
import QualificationsDefault from "./Qualifications-Default.vue";
import TypingField from "./Typing-Field.vue";

export default defineComponent({
  name: "TextField",
  components: { QualificationsDefault, TypingField },
  props: {
    propsDate: String,
  },
  setup() {
    const selectedItem = ref(1);
    let num = ref<number[]>([]);
    let anyNumber = ref(1);
    let themeSelection = ref<string>("dark");
    const date = ref<string | undefined>("");
    const envItems = ["School", "Office"];

    function add() {
      num.value.push(1);
    }

    const removeItems = () => {
      localStorage.removeItem("inputs");
    };

    const showDate = (propsDate: string | undefined) => {
      // console.log(propsDate);
      date.value = propsDate;
      console.log("propsDate = ", date.value);
    };

    const save = (propsDate: string | undefined, text: string) => {
      console.warn("saving...");
    };

    return {
      selectedItem,
      num,
      anyNumber,
      themeSelection,
      date,
      envItems,
      save,
      showDate,
      removeItems,
      add,
    };
  },
});
</script>
