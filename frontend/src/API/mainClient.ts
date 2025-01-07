import { QueryClient } from "@tanstack/react-query";
import axios from "axios";

export const queryClient = new QueryClient({
  defaultOptions: {},
});

export const axiosClient = axios.create({
  baseURL: `http://${location.hostname}:1234`,
});
