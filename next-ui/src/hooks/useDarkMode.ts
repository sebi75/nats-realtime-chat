import { useEffect, useState } from "react";

type SetValueType<T> = (value: T) => void;

const useLocalStorage = <T>(
  key: string,
  initialValue: T
): [T, SetValueType<T>] => {
  const [storedValue, setStoredValue] = useState<T>(() => {
    if (typeof window === "undefined") {
      return initialValue;
    }
    const item = window.localStorage.getItem(key);
    return item ? (JSON.parse(item) as T) : initialValue;
  });

  const setValue: SetValueType<T> = (value) => {
    if (typeof window === "undefined") {
      console.warn("Cannot set local storage during server-side rendering.");
      return;
    }
    setStoredValue(value);
    window.localStorage.setItem(key, JSON.stringify(value));
  };

  return [storedValue, setValue];
};

export const useDarkMode = (): [boolean, SetValueType<boolean>] => {
  const [enabled, setEnabled] = useLocalStorage<boolean>("dark-theme", false);
  const isEnabled = enabled;

  useEffect(() => {
    if (typeof window === "undefined") return;
    const className = "dark";
    const bodyClass = window.document.body.classList;

    isEnabled ? bodyClass.add(className) : bodyClass.remove(className);
  }, [enabled, isEnabled]);

  return [enabled, setEnabled];
};
