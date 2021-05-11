import 'package:flutter/material.dart';

import 'screens/home_screen.dart';

void main() => runApp(MainApp());

class MainApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Moon Base',
      home: HomeScreen(),
    );
  }
}
