<script setup>
import DataTable from "/node_modules/primevue/datatable"
import Column from "/node_modules/primevue/column"
</script>

<template>

  <!-- table -->
  <div class="flex flex-column pt-3 pl-6 pr-6">
    <div class="align-items-center justify-content-center">
      <DataTable :value="products" :dataKey="id" responsiveLayout="scroll">
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
        num: 0,
        }
    },
    created() {
      this.columns = [
        {field: 'no', header: 'No'},
        {field: 'product_name', header: 'Product Name'},
        {field: 'product_value', header: 'Product Value'},
        {field: 'product_description', header: 'Product Description'},
      ];

      this.$watch (
        () => this.$route.params,
        () => {
          this.fetchData()
        },

        { immediate: true }
      )

      this.getNum()

    },
    methods: {
      fetchData() {
        fetch("http://107.102.183.168:8081/v1/api/product/show_products")
          .then(async response => {
            const data = await response.json();
            console.log(data)

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
      getNum() {
        let data = this.products;
        for(var i = 0; i < Object.keys(data).length; i++) {
          num = num + 1;
        };
        console.log(num);
      }
    }
  }

</script>

<style>

</style>
