tailwind.config =  {
    content: [
      "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
      "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
      "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    ],
    theme: {
      colors: {
        "background-dark-blue": "#011F4B",
        "container-dark-blue": "#083158",
        "container-light-blue": "#005B96",
        "select-highlight-blue": "#D0E5F5",
        "hover-highlight-blue": "#B3CDE0",
        "button-highlight-blue": "#819FB5",
        "button-select-blue": "#5F7687",
        "confirm-green": "#2FA71B",
        "confirm-green-hover": "#258016",
        "confirm-green-active": "#1C6110",
        "cancel-red": "#BD342B",
        "cancel-red-hover": "#9E2720",
        "cancel-red-active": "#7A231D",
        "outline-gray": "#868A8f",
        black: "#000000",
        white: "#FFFFFF",
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
  