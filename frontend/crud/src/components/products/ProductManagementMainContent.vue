<script setup>
import DataTable from "/node_modules/primevue/datatable"
import Column from "/node_modules/primevue/column"
</script>

<template>

  <!-- table -->
  <div class="flex flex-column pt-3 pl-6 pr-6">
    <div class="align-items-center justify-content-center">
      <DataTable :value="products" responsiveLayout="scroll">
        <Column v-for="col of columns" :field="col.field" :header="col.header" :key="col.field"></Column>
      </DataTable>
    </div>
  </div>

</template>

<script>
export default {
    data() {
      return {
        columns: null,
        products: null,
        errorMessage: null,
        }
    },
    created() {
      this.columns = [
        {field: 'num', header: 'No'},
        {field: 'product_name', header: 'Product Name'},
        {field: 'product_value', header: 'Product Value'},
        {field: 'product_description', header: 'Product Description'},
      ];

      // async functions
      this.$watch (
        () => this.$route.params,
        () => {
          this.fetchData();
        },

        { immediate: true }
      )

    },
    methods: {
      fetchData() {
        const requestOptions = {
          method: "GET",
          mode: 'cors',
          credentials: 'same-origin',
          cache: 'no-cache'
        };
        fetch("http://107.102.183.168:8081/v1/api/product/show_products", requestOptions)
          .then(async response => {
            const data = await response.json();

            if(!response.ok) {
              const error = (data && data.message) || response.statusText;
              return Promise.reject(error);
            }

            this.products = data;

          })
          .catch(error => {
            this.errorMessage = error;
          });
      },
    }
  }

</script>

<style>

</style>
