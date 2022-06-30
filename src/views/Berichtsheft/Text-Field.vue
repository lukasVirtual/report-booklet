<template>
  <div id="comp">
    <v-container fluid>
      <v-card
        style="width: 900px; border: 2px solid black; overflow-x: hidden"
        tile
        class="mt-5 overflow-y-auto"
        max-width="800px"
      >
        <v-card-title>{{ propsDate }}</v-card-title>
      </v-card>
      <v-card
        style="
          width: 900px;
          height: 400px;
          border: 2px solid black;
          overflow-x: hidden;
        "
        tile
        class="mt-5 overflow-y-auto"
        max-width="800px"
      >
        <v-row dense>
          <v-row>
            <v-col cols="12" md="12">
              <div v-for="j in anyNumber" :key="j">
                <v-card-actions>
                  <v-text-field
                    :auto-grow="true"
                    spellcheck="true"
                    v-model.lazy="text"
                    @change="log"
                  ></v-text-field>
                  <timer-text-field></timer-text-field>
                </v-card-actions>
              </div>
            </v-col>
            <v-spacer></v-spacer>
            <v-col cols="2"> </v-col>
            <v-spacer></v-spacer>
          </v-row>
          <v-col cols="1" class="text-center">
            <div style="background-color: lightblue; height: 320%">
              <div style="height: 5px"></div>
              <v-btn @click="anyNumber++" icon elevation="20"
                ><v-icon color="blue">mdi-plus</v-icon></v-btn
              >
              <div style="height: 10px"></div>
              <v-btn icon elevation="20"
                ><v-icon color="red">mdi-trash-can</v-icon></v-btn
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
import { defineComponent, onMounted, ref, watchEffect } from "vue";
import TimerTextField from "./Timer-TextField.vue";
import QualificationsDefault from "./Qualifications-Default.vue";

export default defineComponent({
  name: "TextField",
  props: {
    propsDate: String,
  },
  setup() {
    const selectedItem = ref(1);
    const items = [{ title: "00:00" }, { title: "00:30" }, { title: "01:00" }];
    let input = ref("00:00");
    let num = ref<number[]>([]);
    let anyNumber = ref(1);
    let themeSelection = ref<string>("dark");
    let text = ref("");

    watchEffect(() => {
      anyNumber.value = Number(localStorage.getItem("inputs"));
    });

    watchEffect(() =>
      localStorage.setItem("inputs", anyNumber.value.toString())
    );
    function add() {
      num.value.push(1);
    }
    const log = () => {
      console.log(text.value);
      localStorage.setItem("va", text.value);
    };

    return {
      selectedItem,
      items,
      num,
      input,
      anyNumber,
      themeSelection,
      text,
      log,
      add,
    };
  },
  components: { TimerTextField, QualificationsDefault },
});
</script>