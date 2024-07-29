import 'package:get/get.dart';
import 'package:movil/pages/page.dart';

class PageNames {
  static const splash = "/";
  static const signIn = "/signin";
  static const dashboard = "/dashboard";
}

// Suggested code may be subject to a license. Learn more: ~LicenseLog:3888769737.
List<GetPage> pages = [
  GetPage(
    name: PageNames.splash,
    page: () => const SplashView(),
    binding: SplashBinding(),
  ),
  GetPage(
    name: PageNames.signIn,
    page: () => const SigninView(),
    binding: SigninBinding(),
  ),
];
