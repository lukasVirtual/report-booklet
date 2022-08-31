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

        <v-card width="550" height="450" rounded>
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
                  color="cyan"
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
                  color="cyan"
                  type="password"
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
                  color="cyan"
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

      <v-dialog v-model="relationDialog" width="600px">
        <template v-slot:activator="{ attrs }">
          <v-btn
            style="margin-left: 10px"
            icon
            v-bind="attrs"
            @click="relationDialog = true"
            ><v-icon>mdi-relation-one-or-many-to-only-one</v-icon></v-btn
          >
        </template>

        <v-card width="550" height="450" rounded>
          <v-card-title class="text-h5 grey lighten-2">
            Assign User to Instructor
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-row>
              <v-col>
                <v-select
                  label="Select Instructor"
                  variant="underlined"
                  color="cyan"
                  :items="instructors"
                ></v-select>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="2">
                <v-btn icon @click="counter++"><v-icon>mdi-plus</v-icon></v-btn>
              </v-col>
              <v-col cols="12" v-for="i in counter" :key="i">
                <v-select
                  label="Select User"
                  variant="underlined"
                  color="cyan"
                  :items="users"
                >
                </v-select>
              </v-col>
            </v-row>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-btn color="red" text @click="relationDialog = false">
              Cancel
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn color="blue" text @click="assignUser"> Assign </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-table
        style="width: 1100px; max-width: 1400px; border-radius: 10px"
        class="mt-5"
      >
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
import { dataService } from "@/handler/dataHandler";
import { defineComponent, onMounted, reactive } from "@vue/runtime-core";
import { inject, ref } from "vue";
import BaseLayout from "../../boilerplate/layouts/Base.vue";
import { useToast } from "vue-toastification";

export default defineComponent({
  name: "AdminDefault",
  components: { BaseLayout },

  setup() {
    const role = ref(inject("role"));
    const toast = useToast();
    let instructors: string[] = [];
    let users: string[] = [];

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
    const relationDialog = ref(false);
    const counter = ref(1);
    const roleItems = ["admin", "instructor", "user"];

    onMounted(async () => {
      const data = await dataService.getAllData();
      for (const elem of data) {
        resultArr.value.push({
          name: elem.Name,
          role: elem.Role,
        });
      }

      resultArr.value.map((value, idx) => {
        if (value.name == "") resultArr.value.splice(idx, 1);
      });

      let getInstructors = data.filter(
        (v: any) => v.Role === "instructor" || v.Role === "user"
      );
      for (const p of getInstructors) {
        if (p.Role === "instructor") instructors.push(p.Name);
        else if (p.Role === "user") users.push(p.Name);
      }
    });

    const addNewUser = async () => {
      if (
        input.name !== "" &&
        input.password.length >= 8 &&
        input.role !== ""
      ) {
        loginService.register(input.name, input.password, input.role);
        resultArr.value.push({ name: input.name, role: input.role });
        input.name = "";
        input.password = "";
        input.role = "";
      } else toast.error("Some fields are Empty or incorrect entered");
      dialog.value = false;
    };

    const deleteElemAt = async (elem: any, idx: number) => {
      if (
        !confirm(
          "You sure you want to delete? Deleted Items will be gone forever"
        )
      )
        return;
      try {
        await dataService.deleteItem(elem.name);
        resultArr.value.splice(idx, 1);
        toast.success("successfully deleted user");
      } catch (e) {
        console.log(e);
      }
      await dataService.getAllData();
    };

    const assignUser = () => {
      toast.success("Successfully assigned User/Users to Instructor");
      relationDialog.value = false;
    };

    return {
      resultArr,
      dialog,
      roleItems,
      input,
      rules,
      role,
      relationDialog,
      instructors,
      users,
      counter,
      assignUser,
      addNewUser,
      deleteElemAt,
    };
  },
});
</script>
