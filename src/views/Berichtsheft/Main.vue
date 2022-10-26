<template>
  <base-layout v-if="$route.path === '/Berichtsheft'">
    <v-main>
      <v-card color="transparent" class="d-flex justify-center mb-6">
        <v-card-actions>
          <v-btn icon @click="switchPageLeft"
            ><v-icon>mdi-arrow-left</v-icon></v-btn
          >
          <v-btn @click="today" style="font-size: 20px">Heute</v-btn>
          <v-btn icon @click="switchPageRight"
            ><v-icon>mdi-arrow-right</v-icon></v-btn
          >

          <v-btn :disabled="toggleStage" @click="submit" color="blue"
            >Submit</v-btn
          >

          <h5 style="color: grey">{{ time }}</h5>
        </v-card-actions>
      </v-card>

      <v-container
        class="d-flex justify-center mb-6"
        v-for="i in daysOfMonth"
        :key="`${i}.${currMonth}.${currYear}`"
        ref="container"
        id="test"
      >
        <text-field :propsDate="`${i}.${currMonth}.${currYear}`"></text-field>
      </v-container>
    </v-main>

    <template v-slot:navIcons>
      <v-btn icon @click="exportAsPdf"
        ><v-icon style="width: 35px" size="22"
          >mdi-file-export-outline</v-icon
        ></v-btn
      >
    </template>
    <v-overlay :model-value="overlay" class="align-center justify-center">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
    <router-view />
  </base-layout>
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
import { dataService } from "@/handler/dataHandler";
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted, ref } from "@vue/runtime-core";
import { useToast } from "vue-toastification";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import TextField, { date } from "./Text-Field.vue";
import { io } from "socket.io-client";

export default defineComponent({
  components: { BaseLayout, TextField },
  name: "MainPage",

  setup() {
    const calendarDate = new Date();
    const role = ref("");
    const toast = useToast();
    const overlay = ref(false);
    const socket = io("http://localhost:7000");

    let daysOfMonth = ref(31);
    let toggleStage = ref(false);
    let currMonth = ref(calendarDate.getMonth() + 1);
    const currYear = ref(calendarDate.getFullYear());
    let time = ref<string | null>(null);

    onMounted(async () => {
      role.value = (await loginService.getUserRole()) as unknown as string;
    });

    const computeDays = () => {
      if (
        currMonth.value === 4 ||
        currMonth.value === 6 ||
        currMonth.value === 9 ||
        currMonth.value === 11
      ) {
        daysOfMonth.value--;
      } else if (currMonth.value !== 2) {
        daysOfMonth.value = 31;
      }

      if (currMonth.value == 2) {
        daysOfMonth.value = daysOfMonth.value - 3;

        if (currYear.value % 4 == 0) daysOfMonth.value++;
        if (currYear.value % 100 == 0) daysOfMonth.value--;
        if (currYear.value % 400 == 0) daysOfMonth.value++;
      }
    };

    const switchPageRight = () => {
      currMonth.value++;
      if (currMonth.value > 12) {
        currMonth.value = 1;
        currYear.value++;
      }
      computeDays();
    };

    const switchPageLeft = async () => {
      if (currYear.value >= calendarDate.getFullYear()) {
        currMonth.value--;
        if (currMonth.value < 1) {
          currMonth.value = 12;
          currYear.value--;
        }
      }
      computeDays();
    };

    const today = () => {
      currMonth.value = calendarDate.getMonth() + 1;
      currYear.value = calendarDate.getFullYear();
      computeDays();
    };

    const submit = async () => {
      overlay.value = true;
      toggleStage.value = !toggleStage.value;
      time.value = `${calendarDate.getHours()}:${calendarDate.getMinutes()}:${calendarDate.getSeconds()}`;
      // console.log(date.value?.split("."));
      const jsonOuput = await dataService.GetJsonMonth(
        date.value?.split(".")?.[1] as string
      );
      try {
        socket.emit("submit", jsonOuput, await loginService.getUser());
        toast.success("successfully submitted");
      } catch (e) {
        toast.error("something went wrong submitting data");
      }
      overlay.value = false;
    };

    const exportAsPdf = async () => {
      overlay.value = true;
      try {
        await dataService.ExportAsPDF();
        toast.success("Succesfully Created PDF");
      } catch (e) {
        toast.error("Error while Exporting");
      }
      overlay.value = false;
    };

    return {
      daysOfMonth,
      currMonth,
      currYear,
      toggleStage,
      time,
      role,
      overlay,
      exportAsPdf,
      submit,
      switchPageRight,
      switchPageLeft,
      today,
    };
  },
});
</script>
