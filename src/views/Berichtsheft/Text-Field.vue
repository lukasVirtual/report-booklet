<template>
  <div id="comp" v-if="$route.path === '/Berichtsheft'">
    <v-container fluid>
      <v-card
        style="width: 900px; border: 2px solid black; overflow-x: hidden"
        tile
        class="mt-5 overflow-y-auto"
      >
        <v-card-title @click="returnDate(propsDate)">{{
          propsDate
        }}</v-card-title>
      </v-card>
      <v-card
        style="
          width: 900px;
          height: 500px;
          border: 1px solid black;
          overflow-x: hidden;
        "
        tile
        class="mt-4 overflow-y-auto"
      >
        <v-row dense>
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
          <v-col cols="1" class="text-center">
            <div
              style="
                background-image: linear-gradient(180deg, lightblue, azure);
                height: 320%;
              "
            >
              <div style="height: 5px"></div>
              <v-btn @click="anyNumber++" icon elevation="20"
                ><v-icon size="25" color="blue">mdi-plus</v-icon></v-btn
              >
              <div style="height: 10px"></div>
              <v-btn icon elevation="20"
                ><v-icon size="25" color="red" @click="removeItems"
                  >mdi-trash-can</v-icon
                ></v-btn
              >
              <div style="height: 10px"></div>

              <qualifications-default></qualifications-default>
            </div>
          </v-col>
        </v-row>
      </v-card>
    </v-container>
  </div>
</template>

<script lang="ts">
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
    let input = ref("00:00");
    let num = ref<number[]>([]);
    let anyNumber = ref(1);
    let themeSelection = ref<string>("dark");

    function add() {
      num.value.push(1);
    }

    const removeItems = () => {
      localStorage.removeItem("inputs");
    };

    const returnDate = (date: string | undefined) => {
      console.log(date);
    };

    return {
      selectedItem,
      num,
      input,
      anyNumber,
      themeSelection,
      returnDate,
      removeItems,
      add,
    };
  },
});
</script>
