const doughnut = {
  type: 'doughnut',
  data: {
    datasets: [],
    labels: []
  },
  options: {
    responsive: true,
    legend: {
      position: 'top',
    },
    title: {
      display: false
    },
    animation: {
      animateScale: true,
      animateRotate: true
    }
  }
};

const timeseries = {
  data: {
    datasets: [{
      label: 'CPU',
      backgroundColor: 'green',
      borderColor: 'green',
      type: 'line',
      pointRadius: 0,
      fill: false,
      lineTension: 0,
      borderWidth: 2,
      data: []
    }, {
      label: 'Memory',
      backgroundColor: 'blue',
      borderColor: 'blue',
      type: 'line',
      pointRadius: 0,
      fill: false,
      lineTension: 0,
      borderWidth: 2,
      data: []
    }],
  },
  options: {
    animation: {
      duration: 0
    },
    scales: {
      xAxes: [{
        type: 'time',
        distribution: 'series',
        offset: true,
        time: {
          unit: 'second'
        }
      }],
      yAxes: [{
        gridLines: {
          drawBorder: false
        },
        scaleLabel: {
          display: false,
          labelString: 'Usage'
        }
      }]
    }
  }
};

function getDoughnutCanvasConfig({ color, initialValue }) {
  const config = Object.assign({}, doughnut);
  config.data = {
    datasets: [{
      data: initialValue,
      backgroundColor: [color, '#efefef']
    }],
    labels: []
  };
  return config;
}

const App = {
  mounted() {
    cpuConfig = getDoughnutCanvasConfig({ color: 'green', initialValue: [0, 100] });
    cpuCanvas = document.querySelector('#cpu').getContext('2d');
    this.cpu = {
      config: cpuConfig,
      chart: new Chart(cpuCanvas, cpuConfig)
    };

    memoryConfig = getDoughnutCanvasConfig({ color: 'blue', initialValue: [0, 100] });
    memoryCanvas = document.querySelector('#memory').getContext('2d');
    this.memory = {
      config: memoryConfig,
      chart: new Chart(memoryCanvas, memoryConfig)
    };

    const dt = moment().valueOf();
    const historyConfig = Object.assign({}, timeseries);
    historyConfig.data.datasets[0].data.push({ t: dt, y: 0 });
    historyConfig.data.datasets[1].data.push({ t: dt, y: 0 }); 
    historyCanvas = document.querySelector('#history').getContext('2d');
    this.history = {
      config: historyConfig,
      chart: new Chart(historyCanvas, historyConfig)
    };
  },
  data() {
    return {
      started: false,
      loading: false,
      url: '',
      interval: 3,
      metrics: {
        runtime: {
          num_goroutine: 0,
          num_gc: 0,
          live_objects: 0
        }
      },
      cpu: {},
      memory: {},
      history: {}
    }
  },
  watch: {
    metrics(val) {
      this.updateCpuChart(val);
      this.updateMemoryChart(val);
      this.updateHistoryChart(val);
    }
  },
  methods: {
    async fetchMetrics() {
      this.loading = true
      try {
        const res = await fetch(this.url);
        const json = await res.json();
        this.metrics = json.data;
        const { interval } = this
        if (interval > 0) {
          setTimeout(this.fetchMetrics, interval * 1000);
        }
      } catch(err) {
        console.warn(err);
      }
      this.loading = false
    },
    start() {
      this.started = true
      this.fetchMetrics();
    },
    updateCpuChart(metrics) {
      const usage = metrics.cpu_usage_percentage;
      const config = this.cpu.config;
      config.data.datasets[0].data = [usage, 100 - usage];
      config.data.labels = [`${usage.toFixed(2)}%`];
      this.cpu.config = config;
      this.cpu.chart.update();
    },
    updateMemoryChart(metrics) {
      const usage = metrics.memory_usage_in_mib;
      const total = metrics.memory_total_in_mib;
      const config = this.memory.config;
      const usagePercentage = (usage * 100) / total;
      config.data.datasets[0].data = [usagePercentage, 100 - usagePercentage];
      config.data.labels = [`${usage} MiB`, `${total} MiB`];
      this.memory.config = config;
      this.memory.chart.update();
    },
    updateHistoryChart(metrics) {
      const memoryTotal = metrics.memory_total_in_mib;
      const memoryUsage = (metrics.memory_usage_in_mib * 100) / memoryTotal;
      const cpuUsage = metrics.cpu_usage_percentage;
      const config = this.history.config;
      const dt = moment().valueOf();
      const maxEntries = 300 / this.interval
      if (maxEntries > config.data.datasets[0].data.length) {
        config.data.datasets[0].data.push({ t: dt, y: cpuUsage });
        config.data.datasets[1].data.push({ t: dt, y: memoryUsage });
      } else {
        config.data.datasets[0].data = [{ t: dt, y: cpuUsage }];
        config.data.datasets[1].data = [{ t: dt, y: memoryUsage }];
      }
      this.history.config = config;
      this.history.chart.update();
    }
  }
};

document.addEventListener('DOMContentLoaded', function() {
  const app = Vue.createApp(App);
  app.component('c-canvas-cpu', { template: '<canvas id="cpu"></canvas>' });
  app.component('c-canvas-memory', { template: '<canvas id="memory"></canvas>' });
  app.component('c-canvas-history', { template: '<canvas id="history"></canvas>' });
  app.component('c-loading', { template: '<div class="lds-hourglass"></div>' });
  app.mount('#app');
  document.querySelector('#app').classList.remove('is-invisible');
});
