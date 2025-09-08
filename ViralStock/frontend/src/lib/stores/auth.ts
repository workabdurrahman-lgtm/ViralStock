import { readable } from 'svelte/store';
import { onAuthStateChanged, type User } from 'firebase/auth';
import { auth } from '$lib/firebase';

// Create a readable store that listens to Firebase auth state changes
export const user = readable<User | null | undefined>(undefined, (set) => {
  const unsubscribe = onAuthStateChanged(
    auth,
    (user) => {
      set(user);
    },
    (error) => {
      console.error('Error in onAuthStateChanged:', error);
      set(null);
    }
  );

  // Return the unsubscribe function to be called when the store is no longer needed
  return () => unsubscribe();
});
