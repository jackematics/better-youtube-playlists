const HISTORY = [];

function add(id) {
  if (HISTORY[HISTORY.length - 1] !== id) {
    HISTORY.push(id);
  }

  if (HISTORY.length > 30) {
    HISTORY.shift();
  }
}

function getPrevious() {
  HISTORY.pop();

  return HISTORY[HISTORY.length - 1];
}

export const History = {
  add,
  getPrevious,
};
