import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:movil/routes/route.dart';

void main() {
  runApp(const Application());
}

class Application extends StatelessWidget {
  const Application({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        useMaterial3: true,
      ),
      getPages: pages,
      initialRoute: RouteNames.splashScreen,
    );
  }
}
