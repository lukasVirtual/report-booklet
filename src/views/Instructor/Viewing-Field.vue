<template>
  <div
    @vnode-mounted="
      cacheData = data;
      propsDate = date;
    "
    v-if="$route.path === '/Instructor'"
  >
    <v-card
      style="
        width: 1100px;
        height: 50px;
        border: 1px solid black;
        overflow-x: hidden;
        border-radius: 10px;
      "
      max-height="500"
      class="mb-2"
    >
      <v-card-title>{{ propsDate }}</v-card-title>
    </v-card>
    <v-card
      style="
        width: 1100px;
        height: 500px;
        border: 1px solid black;
        overflow-x: hidden;
        border-radius: 10px;
      "
      max-height="500"
    >
      <v-row dense>
        <div style="width: 5px"></div>
        <v-row style="max-height: 30px">
          <v-col cols="12" md="12">
            <div v-for="report in reports" :key="report.Input">
              <v-card-actions>
                <typing-field
                  :input="report.Input"
                  :time="report.Time"
                  :id="report.ID"
                ></typing-field>
              </v-card-actions>
            </div>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
      </v-row>
    </v-card>
  </div>
</template>

<script lang="ts">
import { dataService } from "@/handler/dataHandler";
import { defineComponent, onMounted, ref } from "@vue/runtime-core";
import TypingField from "../Berichtsheft/Typing-Field.vue";

export default defineComponent({
  name: "ViewingField",
  components: { TypingField },
  props: {
    data: [Object],
    date: String,
  },

  setup() {
    const cacheData = ref<any>([]);
    const reports = ref<any>([]);
    const propsDate = ref<string | undefined>("");
    onMounted(async () => {
      reports.value = cacheData.value.filter(
        (val: any) => val.Date === propsDate.value
      );
    });

    return { cacheData, propsDate, reports };
  },
});
</script>