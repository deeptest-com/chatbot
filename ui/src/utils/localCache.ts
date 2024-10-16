import { toRaw } from 'vue';
import localforage from 'localforage';
import settings from '@/config/settings';

export const getCache = async (key: string): Promise<any | null> => {
  return await localforage.getItem(key);
};

export const setCache = async (key: string, val: any): Promise<boolean> => {
  try {
    await localforage.setItem(key, toRaw(val));
    return true;
  } catch (err) {
    console.log('setCache err', err)
    return false;
  }
};

export const removeCache = async (key: string): Promise<boolean> => {
  try {
    await localforage.removeItem(settings.siteTokenKey);
    return true;
  } catch (error) {
    return false;
  }
};
