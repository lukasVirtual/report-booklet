<template>
  <v-app>
    <v-main
      v-if="$route.path == '/Register'"
      class="text-center"
      style="
        display: flex;
        align-items: center;
        justify-content: right;
        padding-left: 25%;
      "
    >
      <v-form
        tile
        ref="valid"
        style="
          border: 1px solid black;
          height: 500px;
          width: 800px;
          display: flex;
          justify-content: center;
          align-items: center;
        "
      >
        <v-row justify="center">
          <v-card-title>Register</v-card-title>
          <v-col cols="12">
            <v-text-field
              label="Username"
              type="text"
              style="padding-left: 25%; padding-right: 25%"
              :rules="[rules.required]"
              variant="underlined"
              v-model="vInputField.user"
            ></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field
              label="Password"
              type="password"
              variant="underlined"
              :rules="[rules.required, rules.minLen]"
              style="padding-left: 25%; padding-right: 25%"
              v-model="vInputField.passwrd"
            ></v-text-field>
          </v-col>
          <v-col>
            <router-link :to="{ path: '/Login' }">Back to login</router-link>
          </v-col>
          <v-divider length="800px"></v-divider>
          <v-col>
            <v-card-actions>
              <div class="text-right">
                <!-- <v-card-actions> -->

                <v-btn rounded @click="register">Register</v-btn>
                <v-btn color="red" rounded @click="reset">Reset</v-btn>

                <!-- <v-btn rounded @click="reset">reset</v-btn>
            </v-card-actions> -->
              </div>
            </v-card-actions>
          </v-col>
        </v-row>
      </v-form>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";

export default defineComponent({
  name: "RegisterDefault",

  data() {
    const map: Map<string, string> = new Map([]);
    const vInputField = reactive({
      user: "",
      passwrd: "",
    });

    const rules = {
      required: (value: string) => !!value || "Required",
      minLen: (v: string) => v.length >= 8 || "Min. 8 Characters",
    };

    const register = () => {
      alert("succesfully registerd");
      map.set(vInputField.user, vInputField.passwrd);
      console.error(map);
      for (const [user, password] of map) {
        localStorage.removeItem("registeredUser");
        localStorage.removeItem("registeredUsersPassword");
        localStorage.setItem("registeredUser", user);
        localStorage.setItem("registeredUsersPassword", password);
      }
      this.$router.push({
        name: "Login",
        query: { redirect: "/" },
      });
    };

    const reset = () => {
      const obj = this.$refs.valid as any;
      if (obj !== undefined) obj.reset();
    };
    return { register, rules, reset, vInputField };
  },
});
</script>