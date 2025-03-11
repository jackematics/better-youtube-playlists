tailwind.config = {
  theme: {
    colors: {
      "background-dark-blue": "#011F4B",
      "container-dark-blue": "#083158",
      "container-light-blue": "#005B96",
      "blue-highlight-hover": "#062542",
      "blue-highlight-active": "#66a7e3",
      "warm-green-button": "#027148",
      "warm-green-button-hover": "#034f33",
      "warm-green-button-active": "#023623",
      "green-highlight-hover": "#075236",
      "confirm-green": "#2FA71B",
      "confirm-green-hover": "#258016",
      "confirm-green-active": "#1C6110",
      "cancel-red": "#BD342B",
      "cancel-red-hover": "#9E2720",
      "cancel-red-active": "#7A231D",
      "outline-grey-hover": "#a3a8ad",
      "outline-grey": "#868A8f",
      black: "#000000",
      white: "#FFFFFF",
      "warm-orange": "#cc5c23",
      "warm-orange-hover": "#cc6d3d",
      "orange-highlight": "#ff7e00",
    },
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [],
};
