if (!localStorage.getItem("userId")) {
  localStorage.setItem("userId", crypto.randomUUID());
}
