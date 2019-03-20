<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      fixed
      app
    >
      <v-list dense>
        <v-list-tile
          v-for="menu in menus"
          :key="menu.id"
          @click="transition(menu.link)"
        >
          <v-list-tile-action>
            <v-icon>{{ menu.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-content>
            <v-list-tile-title>{{ menu.name }}</v-list-tile-title>
          </v-list-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>ツール</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            <router-view name="main"/>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="indigo" app>
      <span class="white--text"></span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
  declare var astilectron: any;
  import { Component, Vue } from 'vue-property-decorator';
  @Component
  export default class Home extends Vue {
    private drawer: any = null;
    private menus: Object[] = [
      {
        id: "1",
        icon: "home",
        name: "ホーム",
        link: "/home",
      },
      {
        id: "2",
        icon: "settings",
        name: "エクセル ⇒ TSV",
        link: "/convert_excel_to_tsv",
      }
    ];
    transition(link: string) {
      this.$router.push(link);
    }
  }
</script>