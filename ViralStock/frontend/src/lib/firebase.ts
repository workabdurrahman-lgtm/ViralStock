// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyCgOledye4JkSJPUu4Pwi2NEHxQCogL8Cc",
  authDomain: "viralstock-e04fb.firebaseapp.com",
  projectId: "viralstock-e04fb",
  storageBucket: "viralstock-e04fb.firebasestorage.app",
  messagingSenderId: "100596697154",
  appId: "1:100596697154:web:6bdeb6d1434f3311d085db",
  measurementId: "G-27GDPFKXXD"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export { auth };
