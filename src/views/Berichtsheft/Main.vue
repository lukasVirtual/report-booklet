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
      <v-btn icon @click="saving"
        ><v-icon style="width: 35px" size="22"
          >mdi-content-save-outline</v-icon
        ></v-btn
      >
      <v-btn icon
        ><v-icon style="width: 35px" size="22"
          >mdi-file-export-outline</v-icon
        ></v-btn
      >
    </template>
    <router-view />
  </base-layout>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, ref } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import TextField, { save, date } from "./Text-Field.vue";

export default defineComponent({
  components: { BaseLayout, TextField },
  name: "MainPage",

  setup() {
    const calendarDate = new Date();
    const role = ref(inject("role"));

    let daysOfMonth = ref(31);
    let toggleStage = ref(false);
    let currMonth = ref(calendarDate.getMonth() + 1);
    const currYear = ref(calendarDate.getFullYear());
    let time = ref<string | null>(null);
    const saving = () => {
      save(date.value);
    };

    onMounted(() => {
      window.addEventListener("keydown", (e) => {
        if (e.key == "s" && e.altKey) {
          saving();
          e.preventDefault();
        }
      });
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

      // localStorage.getItem(props.propsDate);
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

    const submit = () => {
      toggleStage.value = !toggleStage.value;

      time.value = `${calendarDate.getHours()}:${calendarDate.getMinutes()}:${calendarDate.getSeconds()}`;

      // localStorage.setItem("stage", JSON.stringify(document.body.innerHTML));
    };

    return {
      daysOfMonth,
      currMonth,
      currYear,
      toggleStage,
      time,
      role,
      saving,
      submit,
      switchPageRight,
      switchPageLeft,
      today,
    };
  },
});
</script>
