module.exports = {
  content: ["./src/**/*.{html,js,vue}"],
  theme: {
    extend: {
      colors: {
        theme: "rgb(var(--color-theme) / <alpha-value>)",
      },
    },
  },
  plugins: [],
}
