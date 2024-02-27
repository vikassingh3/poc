<template>
  <chart-component></chart-component>
  <div>
    <canvas id="myChart" width="400" height="400"></canvas>
  </div>
</template>
<script>

//Chart.js is used for data visualization, rendering bar charts based on the data provided
import Chart from 'chart.js/auto';

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  components: {
    'chart-component': {
      mounted() {
        this.getData();
      },
      methods: {
        getData() {
          fetch('http://localhost:8080/data')
            .then(response => response.json())
            .then(data => {
              this.renderChart(data.data);
            })
            .catch(error => {
              console.error('Error fetching data:', error);
            });
        },
        renderChart(data) {
          const canvas = document.getElementById('myChart');
          if (!canvas) {
            console.error('Canvas element not found');
            return;
          }
          const ctx = canvas.getContext('2d');
          new Chart(ctx, {
            type: 'bar',
            data: {
              labels: ['Data 1', 'Data 2', 'Data 3', 'Data 4', 'Data 5', 'Data 6', 'Data 7', 'Data 8', 'Data 9', 'Data 10'],
              datasets: [{
                label: 'Data',
                data: data.split(',').map(Number),
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1
              }]
            },
          });
        }
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
