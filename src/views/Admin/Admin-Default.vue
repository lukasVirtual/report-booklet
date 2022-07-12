<template>
  <base-layout v-if="$route.name === 'Admin'">
    <v-main style="display: flex; margin: auto">
      <h2 class="font-weight-bold text-center mt-5">Admin Center</h2>
      <v-dialog v-model="dialog" width="600px">
        <template v-slot:activator="{ attrs }">
          <v-btn icon @click="dialog = true" v-bind="attrs"
            ><v-icon>mdi-plus</v-icon></v-btn
          >
        </template>

        <v-card width="550" height="450">
          <v-card-title class="text-h5 grey lighten-2">
            Register new User
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-row>
              <v-col>
                <v-text-field
                  label="Name"
                  variant="underlined"
                  v-model="input.name"
                  :rules="[rules.required]"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-text-field
                  label="Password"
                  variant="underlined"
                  v-model="input.password"
                  :rules="[rules.required, rules.minLen]"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-select
                  :items="roleItems"
                  label="Role"
                  variant="underlined"
                  v-model="input.role"
                  :rules="[rules.required]"
                ></v-select>
              </v-col>
            </v-row>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-btn color="red" text @click="dialog = false"> Cancel </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="blue" text @click="addNewUser"> Register </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-table style="width: 1000px" class="mt-5">
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Role</th>
            <th class="text-left"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(elem, idx) in resultArr" :key="idx">
            <td>{{ elem.name }}</td>
            <td>{{ elem.role }}</td>
            <td class="text-right">
              <v-btn
                v-if="elem.role !== 'admin'"
                @click="deleteElemAt(elem, idx)"
                ><v-icon size="20" color="red">mdi-trash-can</v-icon></v-btn
              >
            </td>
          </tr>
        </tbody>
      </v-table>
    </v-main>
  </base-layout>
</template>

<script lang="ts">
import { loginService } from "@/handler/loginHandler";
import { defineComponent, onMounted, reactive } from "@vue/runtime-core";
import { ref } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";

export default defineComponent({
  name: "AdminDefault",
  components: { BaseLayout },

  setup() {
    const resultArr = ref<[{ name: string; role: string }]>([
      { name: "", role: "" },
    ]);

    const rules = {
      required: (value: string) => !!value || "Required",
      minLen: (v: string) => v.length >= 8 || "Min. 8 Characters",
    };

    const input = reactive({
      name: "",
      password: "",
      role: "",
    });

    const dialog = ref(false);

    const roleItems = ["admin", "instructor", "user"];

    onMounted(async () => {
      const data = await loginService.getAllData();
      for (const elem of data) {
        resultArr.value.push({
          name: elem.Name,
          role: elem.Role,
        });
      }

      resultArr.value.map((value, idx) => {
        if (value.name == "") resultArr.value.splice(idx, 1);
      });
    });

    const addNewUser = async () => {
      if (
        input.name !== "" &&
        input.password.length >= 8 &&
        input.role !== ""
      ) {
        loginService.register(input.name, input.password, input.role);
      } else alert("Some fields are Empty or incorrect entered");
      dialog.value = false;
    };

    const deleteElemAt = async (elem: any, idx: number) => {
      try {
        await loginService.deleteItem(elem.name);
      } catch (e) {
        console.log(e);
      }
    };
    return {
      resultArr,
      dialog,
      roleItems,
      input,
      rules,
      addNewUser,
      deleteElemAt,
    };
  },
});
</script>
