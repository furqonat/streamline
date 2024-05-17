import 'package:get/get.dart';
import 'package:movil/pages/page.dart';

class RouteNames {
  static const splashScreen = "/";
  static const loginScreen = "/login";
  static const dashboardScreen = "/dashboard";
}

// Suggested code may be subject to a license. Learn more: ~LicenseLog:3888769737.
List<GetPage> pages = [
  GetPage(
    name: RouteNames.splashScreen,
    page: () => const SplashView(),
    binding: SplashBinding(),
  ),
  GetPage(
    name: RouteNames.loginScreen,
    page: () => const AuthView(),
    binding: AuthBinding(),
  ),
];
