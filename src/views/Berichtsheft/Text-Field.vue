<template>
  <div
    id="comp"
    v-if="$route.path === '/Berichtsheft'"
    @click="showDate(propsDate)"
  >
    <v-container fluid @vnode-mounted="date = propsDate">
      <v-card
        style="
          width: 1180px;
          border: 1px solid black;
          overflow-x: hidden;
          overflow-y: hidden;
          border-radius: 10px;
        "
        tile
        class="mt-5"
      >
        <v-card-actions style="margin: auto">
          <v-card-title>{{ propsDate }}</v-card-title>
          <!-- <v-btn @click="saveeee(propsDate)">safe</v-btn> -->
          <v-spacer></v-spacer>
          <v-select
            :items="envItems"
            @blur="changeStatus(propsDate)"
            variant="underlined"
            style="max-width: 120px; max-height: 60px"
            v-model="selected"
            :menu-props="{
              closeOnClick: true,
              closeOnContentClick: true,
            }"
          ></v-select>
        </v-card-actions>
      </v-card>
      <div style="display: flex">
        <v-card width="80" class="text-center mt-4">
          <div
            style="
              background-color: #3e5ede;
              height: 500px;
              border-radius: 10px;
            "
          >
            <div style="height: 15px"></div>
            <v-btn @click="add(propsDate)" icon elevation="20"
              ><v-icon size="35" color="blue">mdi-plus</v-icon></v-btn
            >
            <div style="height: 10px"></div>
            <v-btn icon elevation="20" @click="remove = !remove"
              ><v-icon size="30" color="error">mdi-trash-can</v-icon></v-btn
            >
            <div style="height: 10px"></div>

            <qualifications-default
              :dateTextField="propsDate"
              :statusTextField="selected"
            ></qualifications-default>
          </div>
        </v-card>
        <v-card
          id="scrollbar-colored"
          style="
            width: 1100px;
            height: 500px;
            border: 1px solid black;
            overflow-x: hidden;

            border-radius: 10px;
          "
          max-height="500"
          class="mt-4 overflow-y-auto"
        >
          <v-row dense>
            <div style="width: 5px"></div>
            <v-row style="max-height: 30px">
              <v-col cols="12" md="12">
                <div v-for="(j, idx) in out" :key="j">
                  <v-card-actions>
                    <v-btn
                      icon
                      v-if="remove"
                      @click.stop="removeItem(propsDate, idx)"
                      ><v-icon>mdi-close</v-icon></v-btn
                    >
                    <typing-field
                      @vnode-mounted="index = out.length"
                      :rows="j.Rows"
                      :input="j.Input"
                      :time="j.Time"
                      :id="idx"
                      @update="returnIndex(j.Id, propsDate)"
                    ></typing-field>
                  </v-card-actions>
                </div>
              </v-col>
              <v-spacer></v-spacer>
            </v-row>
          </v-row>
        </v-card>
      </div>
    </v-container>
    <!-- <v-overlay
      background-color="primary"
      :model-value="overlay"
      class="align-center justify-center"
    >
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay> -->
  </div>
</template>

<style>
#scrollbar-colored {
  scrollbar-base-color: #000;
  scrollbar-color: grey;
}

::-webkit-scrollbar {
  height: 12px;
  width: 12px;
  background: #191919;
}

::-webkit-scrollbar-thumb {
  background: #4c4c4c;
  -webkit-border-radius: 1ex;
  -webkit-box-shadow: 0px 1px 2px rgba(0, 0, 0, 0.75);
}

::-webkit-scrollbar-corner {
  background: #000;
}
</style>

<script lang="ts">
import { store } from "@/handler/store";
import { dataService } from "@/handler/dataHandler";
import {
  defineComponent,
  nextTick,
  onBeforeMount,
  onMounted,
  onServerPrefetch,
  ref,
  watchEffect,
} from "vue";
import QualificationsDefault from "./Qualifications-Default.vue";
import TypingField from "./Typing-Field.vue";

export const save = (propsDate: string | undefined) => {
  console.warn("saving...");
  store.date = propsDate as string;
  console.log("saving props: ", propsDate);
  console.log(store.input);

  dataService.saveForm(store.input, store.time);
  dataService.saveTextField(propsDate as string);
};

export const date = ref<string | undefined>("");
export const timeStmp = ref("");
export default defineComponent({
  name: "TextField",
  components: { QualificationsDefault, TypingField },
  props: {
    propsDate: String,
  },

  setup() {
    const index = ref<number>(0);
    const selectedItem = ref(1);
    let num = ref<number[]>([]);
    let themeSelection = ref<string>("dark");
    const envItems = ["School", "Office", "Remote", "Holiday", "Sick"];
    const inputField = ref("");
    const out = ref<any | null>(null);
    const remove = ref(false);
    const selected = ref("");
    const status = ref<any | null>(null);
    const overlay = ref(false);

    const add = async (propsDate: string | undefined) => {
      console.log(index.value);
      await dataService.writeJson(
        out.value?.length,
        propsDate as string,
        "",
        "00:00",
        1
      );

      out.value = await dataService.readJson(propsDate as string);
      console.error(out.value.length);
    };

    const showDate = (propsDate: string | undefined) => {
      date.value = propsDate;
      console.log("propsDate = ", date.value);
    };

    onMounted(async () => {
      overlay.value = true;
      await Promise.all([
        dataService.ReadStatus(date.value as string),
        dataService.readJson(date.value as string),
      ]).then((v) => {
        status.value = v[0];
        out.value = v[1];
      });
      if (out.value !== null) {
        if (out.value?.length > 0) {
          console.log(out.value);
          // console.warn(out.value[index.value]);
          if (out.value[index.value]) {
            inputField.value = out.value[index.value].Input;
            timeStmp.value = out.value[index.value].Time;
          }
        }
      }
      if (status.value !== null) {
        selected.value = status.value.Status;
      }
      overlay.value = false;
    });

    const returnIndex = async (idx: number, date: string | undefined) => {
      // console.log(store.id);
      console.log(store.input, store.rows);
      console.warn("Creating...");
      dataService.writeJson(
        store.id as number,
        date as string,
        store.input,
        store.time,
        store.rows
      );

      // store.input = "";
      // store.time = "";
    };

    const removeItem = async (date: string | undefined, idx: number) => {
      console.warn(date, idx);

      await dataService.removeJson(date as string, idx);
      out.value = await dataService.readJson(date as string);

      remove.value = !remove.value;
      console.error(out.value);
    };

    const changeStatus = async (date: string | undefined) => {
      await dataService.WriteStatus(date as string, selected.value);
      status.value = await dataService.ReadStatus(date as string);
      console.log(status.value.Status);
      // if (status !== undefined) {
      //   console.log("status: ", status.Status);
      //   selected.value = status.Status;
      // }
      // console.log(selected.value);
    };

    return {
      selectedItem,
      num,
      themeSelection,
      date,
      envItems,
      inputField,
      index,
      out,
      remove,
      selected,
      status,
      overlay,
      changeStatus,
      removeItem,
      returnIndex,
      save,
      showDate,
      add,
    };
  },
});
</script>
