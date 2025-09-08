<script lang="ts">
  import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth';
  import { auth } from '$lib/firebase';
  import { goto } from '$app/navigation';
  import { user } from '$lib/stores/auth';

  // Redirect to dashboard if user is already logged in
  user.subscribe((currentUser) => {
    if (currentUser) {
      goto('/dashboard');
    }
  });

  async function signInWithGoogle() {
    const provider = new GoogleAuthProvider();
    try {
      await signInWithPopup(auth, provider);
      // The onAuthStateChanged listener in the store will handle the user state,
      // and the subscription above will trigger the redirect.
    } catch (error) {
      console.error('Error during sign-in:', error);
      // Optionally, show an error message to the user
    }
  }
</script>

<div class="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4">
  <div class="w-full max-w-md text-center">
    <h1 class="mb-4 text-4xl font-bold">Login to ViralStock</h1>
    <p class="mb-8 text-gray-600">Please sign in with your Google account to continue.</p>
    <button
      on:click={signInWithGoogle}
      class="inline-flex items-center justify-center rounded-md bg-green-500 px-6 py-3 text-white font-semibold shadow-md transition-colors hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
    >
      <svg class="mr-2 h-5 w-5" aria-hidden="true" focusable="false" data-prefix="fab" data-icon="google" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 488 512"><path fill="currentColor" d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 126 23.4 172.9 61.9l-66.8 66.8c-23.4-22.3-54.8-36.1-91.1-36.1-69.1 0-125.8 56.7-125.8 125.8s56.7 125.8 125.8 125.8c76.3 0 114.8-33.4 119.1-85.8H248V261.8h239.1c1.5 10.9 2.9 22.1 2.9 34z"></path></svg>
      Sign in with Google
    </button>
  </div>
</div>
