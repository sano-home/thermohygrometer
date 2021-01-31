module.exports = {
  // Fix nivo error on running built scripts https://github.com/plouc/nivo/issues/1290#issuecomment-752280413
  webpack: config => {
    config.module.rules.push({
      test: /react-spring/,
      sideEffects: true,
    })

    return config
  },
};
